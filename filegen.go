package main

import (
	"flag"
	"fmt"
	"io/fs"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

const MAX_WORKERS = 1024

var (
	txtTemplate  []string
	fileTemplate []string
	loremIpsum   string
	fileCount    int
	fileSize     int
	rawSize      string
	rndSizeMin   int
	rawRndSize   string
	startDir     string
	showHelp     bool
)

func init() {
	flag.IntVar(&fileCount, "count", 1000, "file count")
	flag.StringVar(&rawSize, "size", "1K", "file size")
	flag.StringVar(&rawRndSize, "random-size-min", "", "random size mode, to enable it, set min size of file > 0")
	flag.StringVar(&startDir, "path", "testdata", "a path for create all data")
	flag.BoolVar(&showHelp, "h", false, "show help")
	flag.BoolVar(&showHelp, "help", false, "show help")

	// templates initialization
	for i := 33; i < 127; i++ { // all ascii chars without space
		txtTemplate = append(txtTemplate, string(rune(i)))
	}
	for i := 48; i < 58; i++ { // digits
		fileTemplate = append(fileTemplate, string(rune(i)))
	}
	for i := 97; i < 103; i++ { // a-f chars
		fileTemplate = append(fileTemplate, string(rune(i)))
	}
}

// Grab the n random chars from the template t and get back concatenated result.
func genRandomString(n int, t []string) string {
	var s []string
	for i := 0; i < n; i++ {
		s = append(s, t[rand.Intn(len(t))])
	}
	return strings.Join(s, "")
}

// Create a file with the content of fileSize random chars.
func createFile(p string) error {
	f, err := os.Create(p)
	defer f.Close()
	if err != nil {
		return err
	}
	content := loremIpsum
	if rndSizeMin > 0 {
		limit := rand.Intn(fileSize - rndSizeMin)
		content = content[:rndSizeMin+limit]
	}
	n, err := f.WriteString(content)
	if n != len(content) {
		return fmt.Errorf("wrong %d bytes written", n)
	}
	if err != nil {
		return err
	}
	return nil
}

// A round robin dirs.
type RRD struct {
	items []string
	idx   int
}

func (r *RRD) Lenght() int {
	return len(r.items)
}

func (r *RRD) Next() string {
	r.idx++
	r.idx = r.idx % len(r.items)
	return r.items[r.idx]
}

func (r *RRD) Add(s ...string) {
	r.items = append(r.items, s...)
}

func NewRRD() *RRD {
	return &RRD{idx: -1}
}

func main() {
	// create rrd
	rrd := NewRRD()

	// parse args and show help (if needed)
	flag.Parse()
	if showHelp {
		flag.PrintDefaults()
		os.Exit(0)
	}

	fileSize, err := Parse(rawSize)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rndSizeMin, _ := Parse(rawRndSize)

	if rndSizeMin > 0 {
		fmt.Printf("random file size mode is enabled,\ncreated files will have size between %d and %d bytes\n", rndSizeMin, fileSize)
	}

	// generate Lorem Ipsum text
	loremIpsum = genRandomString(fileSize, txtTemplate)

	// create a two level dirs, used square root of file count e.g.:
	// for 10000 files will be created 100 dirs with 100 dirs inner of every dir
	wg := &sync.WaitGroup{}
	n := int(math.Sqrt(float64(fileCount)))
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			x := filepath.Join(startDir, genRandomString(6, fileTemplate))
			for j := 0; j < n; j++ {
				y := filepath.Join(x, genRandomString(6, fileTemplate))
				err := os.MkdirAll(y, fs.ModePerm)
				if err != nil {
					panic(err)
				}
				rrd.Add(x, y)
			}
		}()
	}
	wg.Wait()

	// calculate worker count
	var workersCount int
	if fileCount > MAX_WORKERS {
		workersCount = MAX_WORKERS
	} else {
		workersCount = fileCount
	}

	filePathChannel := make(chan string, fileCount)

	workers := &sync.WaitGroup{}
	workers.Add(fileCount)

	// start workers
	for i := 0; i < workersCount; i++ {
		go func() {
			for fpath := range filePathChannel {
				defer workers.Done()
				if err := createFile(fpath); err != nil {
					fmt.Println(fpath, err)
				}
			}
		}()
	}

	// generate tasks
	for i := 0; i < fileCount; i++ {
		fpath := filepath.Join(rrd.Next(), genRandomString(6, fileTemplate)+".txt")
		filePathChannel <- fpath
	}
	close(filePathChannel)
	workers.Wait()
}

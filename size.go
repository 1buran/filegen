package main

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const (
	Byte = 1
	Kilo = 1 << (10 * iota)
	Mega
	Giga
)

var Units = map[string]int{
	"b": Byte,
	"k": Kilo,
	"m": Mega,
	"g": Giga,
}

var re = regexp.MustCompile(`(?i)([\d.]+)\s*([bkmg])b?`)

// Parse size from the given string and return parsed bytes.
func Parse(s string) (int, error) {
	if re.MatchString(s) {
		found := re.FindStringSubmatch(s)

		size, err := strconv.ParseFloat(found[1], 64)
		if err != nil {
			return -1, err
		}

		parsedUnits := strings.ToLower(found[2])
		units, ok := Units[parsedUnits]
		if !ok {
			return -1, fmt.Errorf("invalid units %s", parsedUnits)
		}

		if units == Byte {
			// reject fractional amount of bytes
			_, frac := math.Modf(size)
			if frac > 0 {
				return -1, errors.New("fractional of byte detected")
			}
		}

		return int(size * float64(units)), nil
	}
	return -1, errors.New("invalid input string")
}

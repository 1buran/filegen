package main

import "testing"

func TestParse(t *testing.T) {

	t.Run("valid input", func(t *testing.T) {
		testCases := []struct {
			inputs   []string
			expected int
		}{
			// parse integers set
			{[]string{"1b", "1B", "1 b", "1 B", "1 bytes"}, Byte},
			{[]string{"1K", "1k", "1Kb", "1kilo", "1 KB", "1KB", "1   kilo"}, Kilo},
			{[]string{"1M", "1m", "1Mb", "1mega", "1 MB", "1MB", "1   mega"}, Mega},
			{[]string{"1G", "1g", "1Gb", "1giga", "1 GB", "1GB", "1   giga"}, Giga},

			//parse float numbers set
			{[]string{"1.23 K", "1.23k", "1.23Kb", "1.23kilo", "1.23 KB", "1.23KB", "1.23   kilo"}, 1259},
			{[]string{"1.56M", "1.56m", "1.56Mb", "1.56mega", "1.56 MB", "1.56MB", "1.56   mega"}, 1635778},
			{[]string{"1.345G", "1.345g", "1.345Gb", "1.345giga", "1.345 GB", "1.345GB", "1.345   giga"}, 1444182753},
		}

		for _, c := range testCases {
			for _, input := range c.inputs {
				size, err := Parse(input)
				if err != nil || size != c.expected {
					t.Errorf("error or wrong parsed size,\n\tinput string: %#v, expected/parsed size: %d/%d, err: %s", input, c.expected, size, err)
				}
			}
		}
	})

	t.Run("invalid input", func(t *testing.T) {
		testCases := []struct {
			inputs   []string
			expected int
		}{
			{[]string{"1", "1.5B", "k", "1.12", "1,23", "1_23", "1.aK", "1A", "kilo"}, -1},
		}

		for _, c := range testCases {
			for _, input := range c.inputs {
				size, err := Parse(input)
				if err == nil || size != c.expected {
					t.Errorf("no error raised or wrong parsed size,\n\tinput string: %#v, expected/parsed size: %d/%d", input, c.expected, size)
				}
			}
		}
	})

}

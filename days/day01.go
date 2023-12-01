package days

import (
	"AoC_2023_Go/util"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var mp = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func findDigits(line string) int {
	l := 0
	r := len(line) - 1
	for l < len(line) && !isDigit(line[l]) {
		l++
	}
	for r >= 0 && !isDigit(line[r]) {
		r--
	}
	number, _ := strconv.Atoi(string(line[l]) + string(line[r]))
	return number
}

func getCharOrWord(c byte, s string) string {
	if isDigit(c) {
		return string(c)
	} else if s != "" {
		return mp[s]
	}
	return ""
}

func findWords(line string) int {
	pattern, err := regexp.Compile(`(one|two|three|four|five|six|seven|eight|nine)`)
	if err != nil {
		log.Fatal("Bad Regex...")
	}

	var prefix strings.Builder
	suffix := ""
	n := ""

	for i := 0; i < len(line); i++ {
		prefix.WriteByte(line[i])
		n += getCharOrWord(line[i], pattern.FindString(prefix.String()))
		if n != "" {
			break
		}
	}

	for i := len(line) - 1; i >= 0; i-- {
		suffix = string(line[i]) + suffix
		n += getCharOrWord(line[i], pattern.FindString(suffix))
		if len(n) == 2 {
			break
		}
	}

	number, _ := strconv.Atoi(n)
	return number
}

func Day01() {
	path := "./inputs/day01.txt"
	input := util.ReadInput(path)
	p1_val := 0
	p2_val := 0
	for _, line := range input {
		p1_val += findDigits(line)
		p2_val += findWords(line)
	}

	fmt.Printf("Day 1:\n\tPart 1: The sum of the calibration values is %d.\n", p1_val)
	fmt.Printf("\tPart 2: The sum of the calibration values is now %d.\n", p2_val)
}

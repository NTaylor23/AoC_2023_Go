package days

import (
	"AoC_2023_Go/util"
	"fmt"
	"regexp"
	"strings"
)

var mp2 = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
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
	number := (mp2[string(line[l])] * 10) + mp2[string(line[r])]
	return number
}

func getCharOrWord(c byte, s string) int {
	if isDigit(c) {
		return mp2[string(c)]
	} else if s != "" {
		return mp2[s]
	}
	return 0
}

func findWords(line string, p regexp.Regexp) int {
	var prefix strings.Builder
	suffix := ""

	l := 0
	r := len(line) - 1

	number := 0
	for l < len(line) {
		prefix.WriteByte(line[l])
		number += getCharOrWord(line[l], p.FindString(prefix.String())) * 10
		if number != 0 {
			break
		}
		l++
	}

	for r >= 0 {
		suffix = string(line[r]) + suffix
		secondDigit := getCharOrWord(line[r], p.FindString(suffix))
		if secondDigit != 0 {
			number += secondDigit
			break
		}
		r--
	}

	return number
}

func Day01() {
	path := "./inputs/day01.txt"
	input := util.ReadInput(path)
	p1_val := 0
	p2_val := 0
	pattern, _ := regexp.Compile(`(one|two|three|four|five|six|seven|eight|nine)`)
	for _, line := range input {
		p1_val += findDigits(line)
		p2_val += findWords(line, *pattern)
	}

	fmt.Printf("Day 1:\n\tPart 1: The sum of the calibration values is %d.\n", p1_val)
	fmt.Printf("\tPart 2: The sum of the calibration values is now %d.\n", p2_val)
}

package days

import (
	"AoC_2023_Go/util"
	"fmt"
	"strconv"
	"strings"
)

func parseLine(line string) (bool, int) {
	maxVals := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	data := strings.Split(line, ": ")[1]
	for _, group := range strings.Split(data, "; ") {
		for _, colorCount := range strings.Split(group, ", ") {
			s := strings.Split(colorCount, " ")
			amt, _ := strconv.Atoi(s[0])
			color := s[1]
			maxVals[color] = max(amt, maxVals[color])
		}
	}
	isValid := maxVals["red"] <= 12 &&
		maxVals["green"] <= 13 &&
		maxVals["blue"] <= 14

	return isValid, (maxVals["red"] * maxVals["green"] * maxVals["blue"])
}

func Day02() {
	path := "./inputs/day02.txt"
	input := util.ReadInput(path)
	p1_val, p2_val := 0, 0
	for i, line := range input {
		isValid, power := parseLine(line)
		if isValid {
			p1_val += (i + 1)
		}
		p2_val += power
	}

	fmt.Printf("Day 2:\n\tPart 1: The sum of the the IDs of valid games is %d.\n", p1_val)
	fmt.Printf("\tPart 2: The sum of the powers of these games is %d.\n", p2_val)
}

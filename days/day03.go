package days

import (
	"AoC_2023_Go/util"
	"fmt"
	"regexp"
	"strconv"
)

func Day03() {
	path := "./inputs/day03.txt"
	input := util.ReadInput(path)

	numberPattern := regexp.MustCompile(`(\d+)`)      // any number
	symbolPattern := regexp.MustCompile(`([^\.0-9])`) // any symbol
	gearPattern := regexp.MustCompile(`(\*)`)         // any asterisk
	numberPositions := make([][][]int, len(input))

	p1_val, p2_val := 0, 0

	var checkPerimiter = func(row int, cStart int, cEnd int) bool {
		for i := max(row-1, 0); i <= min(row+1, len(input)-1); i++ {
			if symbolPattern.MatchString(input[i][max(cStart-1, 0):min(cEnd+1, len(input[0]))]) {
				return true
			}
		}
		return false
	}

	var checkGear = func(row int, col int) int {
		adjacents := make([]string, 0)
		for i := max(row-1, 0); i <= min(row+1, len(input)-1); i++ {
			for _, numPos := range numberPositions[i] {
				if max(numPos[0]-1, 0) <= col && col <= numPos[1] {
					adjacents = append(adjacents, input[i][numPos[0]:numPos[1]])
				}
			}

		}
		if len(adjacents) == 2 {
			n1, _ := strconv.Atoi(adjacents[0])
			n2, _ := strconv.Atoi(adjacents[1])
			return n1 * n2
		}

		return 0
	}

	for i, line := range input {
		numberPositions[i] = append(numberPositions[i], numberPattern.FindAllStringIndex(line, -1)...)
	}

	for i, row := range numberPositions {
		for _, pos := range row {
			start, end := pos[0], pos[1]
			if checkPerimiter(i, start, end) {
				val, _ := strconv.Atoi(input[i][start:end])
				p1_val += val
			}
		}
		for _, gear := range gearPattern.FindAllStringIndex(input[i], -1) {
			p2_val += checkGear(i, gear[0])
		}
	}
	fmt.Printf("Day 3:\n\tPart 1: The sum of all part numbers is %d.\n", p1_val)
	fmt.Printf("\tPart 2: The sum of all gear ratios is %d.\n", p2_val)
}

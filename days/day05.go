package days

import (
	"AoC_2023_Go/util"
	"fmt"
	"math"
	"regexp"
)

func parseMaps(input []string, pattern *regexp.Regexp) [][][]int {
	input = append(input, "")
	maps := make([][][]int, 0)
	current := make([][]int, 0)

	for _, line := range input {
		if line == "" {
			maps = append(maps, current)
			current = nil
		} else if isDigit(line[0]) {
			current = append(current, util.AtoiIter(pattern.FindAllString(line, -1)))
		}
	}
	return maps
}

func findPathInMap(val int, mp [][]int) int {
	for _, row := range mp {
		dest, source, rng := row[0], row[1], row[2]
		if val >= source && val <= source+rng {
			return dest + (val - source)
		}
	}
	return val
}

func traverse(seed int, maps [][][]int) int {
	for i := 0; i < len(maps); i++ {
		seed = findPathInMap(seed, maps[i])
	}
	return seed
}

func Day05() {
	path := "./inputs/day05.txt"
	input := append(util.ReadInput(path), "")

	p1_val, p2_val := 0, 0

	nPat := regexp.MustCompile(`(\d+)`)
	seeds := util.AtoiIter(nPat.FindAllString(input[0], -1))
	minLoc := int(math.Inf(1))

	maps := parseMaps(input[3:], nPat)
	for _, seed := range seeds {
		minLoc = min(minLoc, traverse(seed, maps))
	}
	p1_val = minLoc
	minLoc = int(math.Inf(1))
	for i := 0; i < len(seeds); i += 2 {
		start, end := seeds[i], seeds[i] + seeds[i + 1]
		fmt.Printf("start: %v\n", start)
		for start <= end {
			minLoc = min(minLoc, traverse(start, maps))
			start++
		}
	}
	p2_val = minLoc

	fmt.Printf("Day 5:\n\tPart 1: The lowest location number is %d.\n", p1_val)
	fmt.Printf("\tPart 2: The lowest location number in any range is %d.\n", p2_val)
}

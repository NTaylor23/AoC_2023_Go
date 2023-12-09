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

	p1_val := 0

	nPat := regexp.MustCompile(`(\d+)`)
	seeds := util.AtoiIter(nPat.FindAllString(input[0], -1))
	minLoc := int(math.Inf(1))

	maps, pairs := parseMaps(input[3:], nPat), make([][]int, 0)

	for i := 0; i < len(seeds); i += 2 {
		pairs = append(pairs, []int{seeds[i], seeds[i] + seeds[i+1]})
	}

	for _, seed := range seeds {
		minLoc = min(minLoc, traverse(seed, maps))
	}
	p1_val = minLoc

	for _, mp := range maps {
		new := make([][]int, 0)
		for len(pairs) > 0 {
			start, end := pairs[len(pairs)-1][0], pairs[len(pairs)-1][1]
			pairs = pairs[:len(pairs)-1]
			flag := false

			for _, m := range mp {
				dest, source, rng := m[0], m[1], m[2]
				overlap_start := max(start, source)
				overlap_end := min(end, source+rng)

				if overlap_start < overlap_end {
					new = append(new, []int{overlap_start - source + dest, overlap_end - source + dest})

					if overlap_start > start {
						pairs = append(pairs, []int{start, overlap_start})
					}

					if end > overlap_end {
						pairs = append(pairs, []int{overlap_end, end})
					}

					flag = true
					break
				}
			}
			if !flag {
				new = append(new, []int{start, end})
			}
		}
		pairs = new
	}
	p2_val := pairs[0][0]

	for _, p := range pairs {
		p2_val = min(p2_val, min(p[0], p[1]))
	}

	fmt.Printf("Day 5:\n\tPart 1: The lowest location number is %d.\n", p1_val)
	fmt.Printf("\tPart 2: The lowest location number in any range is %d.\n", p2_val)
}

package days

import (
	"AoC_2023_Go/util"
	"fmt"
	"strings"

	"github.com/hashicorp/go-set/v2"
)

var STARTS_AT_IDX int = 10

func Day04() {
	path := "./inputs/day04.txt"
	input := util.ReadInput(path)
	cache := make([]int, len(input))
	p1_val := 0

	for card, line := range input {
		cache[card]++
		s := strings.Split(line, " | ")
		matches := set.From[int](
			util.AtoiIter(strings.Fields(s[1]))).Intersect(
			set.From[int](util.AtoiIter(strings.Fields(s[0][STARTS_AT_IDX:]))),
		).Size()
		if matches > 0 {
			p1_val += 1 << (matches - 1)
		}
		for i := card + 1; i < card+matches+1; i++ {
			cache[i] += cache[card]
		}
	}
	fmt.Printf("Day 4:\n\tPart 1: I have %d points in total.\n", p1_val)
	fmt.Printf("\tPart 2: I have %d scratchcards in total.\n", util.Sum(cache))
}

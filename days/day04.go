package days

import (
	"AoC_2023_Go/util"
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/hashicorp/go-set/v2"
)

// the demo input is different from the main input,
// and the actual numbers start at a different index (10 vs 7)
var STARTS_AT_IDX int = 10

func Day04() {
	path := "./inputs/day04.txt"
	input := util.ReadInput(path)

	nPat := regexp.MustCompile(`(\d+)`)
	cache := make([]int, len(input))
	p1_val := 0
	
	for card, line := range input {
		cache[card]++
		s := strings.Split(line, " | ")
		matches := set.From[string](
			nPat.FindAllString(s[1], -1)).Intersect(
				set.From[string](nPat.FindAllString(s[0][STARTS_AT_IDX:], -1)),
		).Size()
		p1_val += int(1 * math.Pow(2, float64(matches - 1)))
		for i := card + 1; i < card+matches+1; i++ {
			cache[i] += cache[card]
		}
	}
	fmt.Printf("Day 4:\n\tPart 1: I have %d points in total.\n", p1_val)
	fmt.Printf("\tPart 2: I have %d scratchcards in total.\n", util.Sum(cache))
}

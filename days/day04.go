package days

import (
	"AoC_2023_Go/util"
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/hashicorp/go-set/v2"
)

func checkNumbers(myNumbers []int, winningNumbers *set.Set[int]) (int, int) {
	score := -1
	matches := 0
	for _, n := range myNumbers {
		if winningNumbers.Contains(n) {
			matches++
			if score > 0 {
				score *= 2
			}
			score = util.AbsVal(score)
		}
	}
	return max(score, 0), matches
}

func Day04() {
	path := "./inputs/day04.txt"
	input := util.ReadInput(path)
	nPat := regexp.MustCompile(`(\d+)`)

	cache := make([]int, len(input))
	for i := 0; i < len(input); i++ {
		cache[i] = 1
	}

	p1_val := 0

	for card, line := range input {
		s := strings.Split(line, " | ")
		l := set.From[string](
			nPat.FindAllString(s[1], -1)).Intersect(
				set.From[string](nPat.FindAllString(s[0][7:], -1)),
		).Size()
		mine := util.AtoiIter(nPat.FindAllString(s[1], -1))
		wins := set.From[int](util.AtoiIter(nPat.FindAllString(s[0][10:], -1)))
		score, matches := checkNumbers(mine, wins)
		sc := int(1 * math.Pow(2, float64(l - 1)))
		fmt.Printf("sc: %v\n", sc)
		p1_val += score
		for i := card + 1; i < card+matches+1; i++ {
			cache[i] += cache[card]
		}
	}

	p2_val := util.Sum(cache)

	fmt.Printf("Day 4:\n\tPart 1: I have %d points in total.\n", p1_val)
	fmt.Printf("\tPart 2: I have %d scratchcards in total.\n", p2_val)
}

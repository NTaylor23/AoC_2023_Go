package days

import (
	"AoC_2023_Go/util"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/hashicorp/go-set/v2"
	"golang.org/x/exp/maps"
)

var NUM_TYPES int = 7

var strength map[rune]int = map[rune]int{
	'2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8,
	'9': 9, 'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14,
}
var alphabet map[rune]byte = map[rune]byte{
	'2': 'B', '3': 'C', '4': 'D', '5': 'E', '6': 'F', '7': 'G', '8': 'H',
	'9': 'I', 'T': 'J', 'J': 'K', 'Q': 'L', 'K': 'M', 'A': 'N',
}

const (
	HIGHCARD = iota + 1
	ONEPAIR
	TWOPAIR
	THREEOAK
	FULLHOUSE
	FOUROAK
	FIVEOAK
)

type Hand struct {
	cards string
	bid   int
	hType int
}

func countCards(bytes []byte) []int {
	mp := make(map[byte]int)
	for _, b := range bytes {
		mp[b]++
	}
	return maps.Values(mp)
}

func replaceJ(s string) string {
	counts := make(map[rune]int)
	var maxChar rune
	maxCount := 0

	for _, char := range s {
		if char == 'J' {
			continue
		}
		counts[char]++
		if counts[char] > maxCount {
			maxCount = counts[char]
			maxChar = char
		}
	}
	return strings.ReplaceAll(s, "J", string(maxChar))
}

func getType(s string, part2 bool) int {
	if part2 {
		s = replaceJ(s)
	}
	to_bytes := []byte(s)
	st := set.From[byte]([]byte(s))
	switch st.Size() {
	case 1:
		return FIVEOAK
	case 2:
		if slices.Min(countCards(to_bytes)) == 2 {
			return FULLHOUSE
		}
		return FOUROAK
	case 3:
		if slices.Max(countCards(to_bytes)) == 3 {
			return THREEOAK
		}
		return TWOPAIR
	case 4:
		return ONEPAIR
	default:
		return HIGHCARD
	}
}

func toAlphabet(s string) string {
	bytes := []byte(s)
	for i, c := range s {
		bytes[i] = alphabet[c]
	}
	return string(bytes)
}

func sortLex(hands [][]Hand) [][]Hand {
	for i := 0; i < NUM_TYPES; i++ {
		sort.Slice(hands[i], func(a, b int) bool {
			return strings.Compare(
				toAlphabet(hands[i][a].cards),
				toAlphabet(hands[i][b].cards)) < 0
		})
	}
	return hands
}

func orderByType(input []string, part2 bool) [][]Hand {
	hands := make([][]Hand, NUM_TYPES)
	for _, hand := range input {
		sp := strings.Split(hand, " ")
		cards := sp[0]
		bid, _ := strconv.Atoi(sp[1])
		t := getType(cards, part2)
		hands[t-1] = append(hands[t-1], Hand{cards, bid, t})
	}
	return hands
}

func solveWinnings(hands [][]Hand) int {
	val := 0
	index := 1
	hands = sortLex(hands)
	for i := 0; i < NUM_TYPES; i++ {
		for j := 0; j < len(hands[i]); j++ {
			val += hands[i][j].bid * index
			index += 1
		}
	}
	return val
}

func Day07() {
	path := "./inputs/day07.txt"
	input := util.ReadInput(path)

	p1_ordering := orderByType(input, false)
	p2_ordering := orderByType(input, true)

	p1_val := solveWinnings(p1_ordering)
	strength['J'] = 1
	alphabet['J'] = 'A'
	p2_val := solveWinnings(p2_ordering)

	fmt.Printf("Day 7:\n\tPart 1: The total winnings are %d.\n", p1_val)
	fmt.Printf("\tPart 2: The new total winnings are %d.\n", p2_val)
}

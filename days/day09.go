package days

import (
	"AoC_2023_Go/util"
	"fmt"
	"strings"
)

func getHistory(nums []int) (int, int) {
	lastElements := []int{nums[len(nums)-1]}
	firstElements := []int{nums[0]}

	l, r := 0, len(nums)-1

	for nums[0] != 0 || nums[r] != 0 {
		for l < r {
			nums[l] = nums[l+1] - nums[l]
			l++
		}
		lastElements = append(lastElements, nums[r-1])
		firstElements = append(firstElements, nums[0])
		r--
		l = 0
	}

	for i := len(firstElements) - 2; i >= 0; i-- {
		firstElements[i] -= firstElements[i+1]
	}

	return util.Sum(lastElements), firstElements[0]
}

func Day09() {
	path := "./inputs/day09.txt"
	input := util.ReadInput(path)
	p1_val, p2_val := 0, 0
	for _, line := range input {
		nums := util.AtoiIter(strings.Split(line, " "))
		last, first := getHistory(nums)
		p1_val += last
		p2_val += first
	}

	fmt.Printf("Day 9:\n\tPart 1: The sum of these extrapolated values is %d.\n", p1_val)
	fmt.Printf("\tPart 2: The new sum of these extrapolated values is %d.\n", p2_val)
}

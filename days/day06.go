package days

import (
	"AoC_2023_Go/util"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func solve(t int, r int) int {
	a := float64(-1)
	b := float64(t)
	c := float64(-(r + 1))
	pos_x := (-b + math.Sqrt(b*b-4*a*c)) / (2 * a)
	neg_x := (-b - math.Sqrt(b*b-4*a*c)) / (2 * a)
	return int(math.Floor(neg_x) - math.Ceil(pos_x) + 1)
}

func Day06() {
	path := "./inputs/day06.txt"
	input := append(util.ReadInput(path), "")

	time, dist := util.IntsInString(input[0]), util.IntsInString(input[1])
	cTime, _ := strconv.Atoi(strings.Join(strings.Fields(input[0])[1:], ""))
	cDist, _ := strconv.Atoi(strings.Join(strings.Fields(input[1])[1:], ""))

	p1_val, p2_val := 1, solve(cTime, cDist)

	for i, t := range time {
		p1_val *= solve(t, dist[i])
	}

	fmt.Printf("Day 6:\n\tPart 1: This race can be beaten %d ways.\n", p1_val)
	fmt.Printf("\tPart 2: This much longer race can be beaten %d ways.\n", p2_val)
}

package days

import (
	"AoC_2023_Go/util"
	"fmt"
	"strconv"
	"strings"
	"math"
)

func quadratic(t, d float64) int {
    discriminant := t*t - 4*d
    if discriminant < 0 {
        return 0
    }
    return int(math.Floor((t - math.Sqrt(discriminant)) / 2))
}

func getRecordTimes(t int, d int) int {
	res := 0
	for i := quadratic(float64(t), float64(d)); i < t; i++ {
		dist := i * (t - i)
		if dist > d {
			res++
		} else if dist < d && res > 1 {
			return res
		}
	}
	return res
}

func Day06() {
	path := "./inputs/day06.txt"
	input := append(util.ReadInput(path), "")

	time, dist := util.IntsInString(input[0]), util.IntsInString(input[1])
	cTime, _ := strconv.Atoi(strings.Join(strings.Fields(input[0])[1:], ""))
	cDist, _ := strconv.Atoi(strings.Join(strings.Fields(input[1])[1:], ""))
	
	p1_val, p2_val := 1, getRecordTimes(cTime, cDist)

	for i, t := range time {
		p1_val *= getRecordTimes(t, dist[i])
	}
	
	fmt.Printf("Day 6:\n\tPart 1: This race can be beaten %d ways.\n", p1_val)
	fmt.Printf("\tPart 2: This much longer race can be beaten %d ways.\n", p2_val)
}

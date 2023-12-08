package days

import (
	"AoC_2023_Go/util"
	"fmt"
	"strings"

	"github.com/hashicorp/go-set/v2"
)

var adj map[string][]string = map[string][]string{}

func processLinear(directions string, curr string, end *set.Set[string]) int {
	index := 0
	steps := 0
	sz := len(directions)
	for !end.Contains(curr) {
		step := directions[index]
		if step == 'L' {
			curr = adj[curr][0]
		} else {
			curr = adj[curr][1]
		}
		index = (index + 1) % sz
		steps++
	}
	return steps
}

func GCD(a, b int) int {
    for b != 0 {
        t := b
        b = a % b
        a = t
    }
    return a
}

func LCM(a, b int) int {
    return a * b / GCD(a, b)
}

func processBidirectional(a []string, directions string, end *set.Set[string]) int {
	result := 1
	for _, start := range a {
		result = LCM(result, processLinear(directions, start, end))
	}
	return result
}

func Day08() {
	path := "./inputs/day08.txt"
	input := util.ReadInput(path)

	startsWithA, startsWithZ := make([]string, 0), make([]string, 0)
	for _, line := range input[2:] {
		sp := strings.Split(line, " = ")
		node := sp[0]
		dests := strings.Split(sp[1], ", ")
		adj[node] = []string{dests[0][1:], dests[1][:len(dests[1])-1]}

		if node[2] == 'A' {
			startsWithA = append(startsWithA, node)
		} else if node[2] == 'Z' {
			startsWithZ = append(startsWithZ, node)
		}
	}

	p1_val := processLinear(input[0], "AAA", set.From([]string{"ZZZ"}))
	p2_val := processBidirectional(startsWithA, input[0], set.From[string](startsWithZ))

	fmt.Printf("Day 8:\n\tPart 1: We need to take %d steps to reach ZZZ.\n", p1_val)
	fmt.Printf("\tPart 2: It takes %d steps before we're at nodes ending only with Z.\n", p2_val)
}

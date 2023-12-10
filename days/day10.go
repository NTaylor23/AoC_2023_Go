package days

import (
	"AoC_2023_Go/util"
	"fmt"
	"slices"
	"strings"

	"github.com/golang-collections/collections/stack"
	"github.com/hashicorp/go-set/v2"
)

type Pos struct {
	y, x int
}

var look map[rune][][]int = map[rune][][]int{
	'|': {{1, 0}, {-1, 0}},
	'-': {{0, 1}, {0, -1}},
	'L': {{-1, 0}, {0, 1}},
	'J': {{-1, 0}, {0, -1}},
	'7': {{1, 0}, {0, -1}},
	'F': {{1, 0}, {0, 1}},
}

func inBounds(y, x, m, n int) bool {
	return (y >= 0 && y < m) && (x >= 0 && x < n)
}

func travel(grid [][]int, input []string, pos Pos, start Pos, seen *set.Set[Pos], m, n int) int {
	stack, dist := stack.New(), 1
	stack.Push(pos)

	for stack.Len() > 0 {
		curr := stack.Pop().(Pos)
		y, x := curr.y, curr.x
		grid[y][x] = 1
		dist++
		seen.Insert(curr)
		for _, adj := range look[rune(input[y][x])] {
			yy, xx := y+adj[0], x+adj[1]
			newPos := Pos{yy, xx}
			if input[yy][xx] != '.' && inBounds(yy, xx, m, n) && !seen.Contains(newPos) {
				stack.Push(Pos{yy, xx})
			}
		}
	}
	return dist / 2
}

func Day10() {
	path := "./inputs/day10.txt"
	input := util.ReadInput(path)
	m, n := len(input), len(input[0])

	grid := make([][]int, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]int, n)
	}

	p := Pos{0, 0}

	for i, v := range input {
		j := strings.Index(v, "S")
		if j > -1 {
			p.y, p.x = i, j
			grid[i][j] = 1
			break
		}
	}

	p1_val := 0
	seen := set.From[Pos]([]Pos{p})
	for _, adj := range [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
		if seen.Size() > (m*n)/2 {
			break
		}
		if inBounds(p.y+adj[0], p.x+adj[1], m, n) {
			new := Pos{p.y + adj[0], p.x + adj[1]}
			p1_val = max(p1_val, travel(grid, input, new, p, seen, m, n))
		}
	}

	p2_val := 0
	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			if seen.Contains(Pos{y, x}) {
				continue
			}

			countIntersect := 0
			ty, tx := y, x
			for ty < m && tx < n {
				tmp := input[ty][tx]
				if !slices.Contains([]byte{'L', '7'}, tmp) && seen.Contains(Pos{ty, tx}) {
					countIntersect++
				}
				ty++
				tx++
			}
			p2_val += countIntersect % 2
		}
	}

	fmt.Printf("Day 10:\n\tPart 1: The farthest starting position is %d steps away.\n", p1_val)
	fmt.Printf("\tPart 2: There are %d tiles enclosed by the loop.\n", p2_val)
}

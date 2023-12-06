package util

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

/// INPUT METHODS

func ReadInput(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		workdir, _ := os.Getwd()
		log.Fatalf("Bad file path: Could not read file in directory %s -> %s", workdir, err)
	}
	defer f.Close()

	res := make([]string, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		res = append(res, strings.TrimRight(line, "\n"))
	}
	return res
}

/// NUMERIC METHODS

func Sum[T constraints.Ordered](numbers []T) T {
	var total T
    for _, number := range numbers {
        total += number
    }
    return total
}

func AtoiIter(numbers []string) []int {
	res := make([]int, len(numbers))
	for i, n := range numbers {
		conv, err := strconv.Atoi(n)
		if err != nil {
			log.Fatalf("could not convert %s to integer", n)
		}
		res[i] = conv
	}
	return res
}

func AbsVal(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func IntsInString(s string) []int {
	nPat := regexp.MustCompile(`(\d+)`)
	return AtoiIter(nPat.FindAllString(s, -1))
}
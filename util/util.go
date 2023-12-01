package util

import (
	"bufio"
    "log"
    "os"
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
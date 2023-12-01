package main

import (
	"AoC_2023_Go/days"
	"fmt"
	"time"
)

type fn func()

func timer(f fn) {
	start := time.Now()
	f()
	duration := time.Since(start)
	fmt.Println("	Completed in:", duration)
}

func main() {
	timer(days.Day01)
}
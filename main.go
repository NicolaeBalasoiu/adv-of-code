package main

import (
	y2020 "aoc/y2020"
	"fmt"
	"runtime"
	"time"
)

func main () {
	var memStats runtime.MemStats
	start := time.Now()

	runtime.ReadMemStats(&memStats)
	fmt.Printf("Memory Usage Before: Alloc = %v KB\n", memStats.Alloc/1024)

	y2020.Day7()

	runtime.ReadMemStats(&memStats)
	fmt.Printf("Memory Usage After: Alloc = %v KB\n", memStats.Alloc/1024)
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
}
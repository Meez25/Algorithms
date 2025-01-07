package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	n, m := 18, 18
	cache := make(map[string]int, n+m)
	result := gridTraveler(n, m, cache)
	fmt.Println(result, "in", time.Since(start))
}

func gridTraveler(n, m int, cache map[string]int) int {
	key := fmt.Sprintf("%d,%d", n, m)
	if val, ok := cache[key]; ok {
		return val
	}
	if n == 0 || m == 0 {
		return 0
	}
	if n == 1 && m == 1 {
		return 1
	}

	// Two ways to travel, right and down. This effectively decreases the size of the grid
	cache[key] = gridTraveler(n-1, m, cache) + gridTraveler(n, m-1, cache)
	return cache[key]
}

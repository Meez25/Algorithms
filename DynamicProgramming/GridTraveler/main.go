package main

import (
	"fmt"
	"time"
)

type Grid struct {
	n int
	m int
}

func main() {
	start := time.Now()
	n, m := 18, 18
	startGrid := Grid{n, m}
	cache := make(map[Grid]int, n*m)
	result := gridTraveler(startGrid, cache)
	fmt.Println(result, "in", time.Since(start))
}

func gridTraveler(grid Grid, cache map[Grid]int) int {
	if val, ok := cache[grid]; ok {
		return val
	}
	if grid.n == 0 || grid.m == 0 {
		return 0
	}
	if grid.n == 1 && grid.m == 1 {
		return 1
	}

	// Two ways to travel, right and down. This effectively decreases the size of the grid
	downGrid := Grid{n: grid.n - 1, m: grid.m}
	rightGrid := Grid{n: grid.n, m: grid.m - 1}
	cache[grid] = gridTraveler(downGrid, cache) + gridTraveler(rightGrid, cache)
	return cache[grid]
}

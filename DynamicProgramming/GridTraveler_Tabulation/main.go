package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	x, y := 18, 18
	result := gridTraveler(x, y)
	fmt.Println(result, "in", time.Since(start))
}

func gridTraveler(n, m int) int {
	table := make([][]int, m+1)
	for i := range table {
		row := make([]int, n+1)
		table[i] = row
	}
	table[1][1] = 1

	for y := range table {
		for x := range table[y] {
			if y < len(table)-1 {
				table[y+1][x] += table[y][x]
			}
			if x < len(table[0])-1 {
				table[y][x+1] += table[y][x]
			}
		}
	}

	return table[m][n]
}

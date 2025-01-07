package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	toCompute := 50
	result := fib(toCompute)
	fmt.Println(result, "in", time.Since(start))
}

func fib(n int) int {
	table := make([]int, n+3)
	for i := range table {
		table[i] = 0
	}
	table[1] = 1

	for i := 0; i <= n; i++ {
		table[i+1] += table[i]
		table[i+2] += table[i]
	}
	return table[n]
}

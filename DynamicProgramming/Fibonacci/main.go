package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	toCompute := 500
	cache := make(map[int]int, toCompute)
	result := fib(toCompute, cache)
	fmt.Println(result, "in", time.Since(start))
}

func fib(n int, cache map[int]int) int {
	if val, ok := cache[n]; ok {
		return val
	}
	if n <= 2 {
		return 1
	}
	cache[n] = fib(n-1, cache) + fib(n-2, cache)
	return cache[n]
}

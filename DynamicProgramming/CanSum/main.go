package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	target := 300
	listOfAvailableNumbers := []int{14, 7}
	cache := make(map[int]bool, target)
	result := canSum(target, listOfAvailableNumbers, cache)
	fmt.Println(result, "in", time.Since(start))
}

func canSum(target int, numbers []int, cache map[int]bool) bool {
	if len(numbers) == 0 {
		return false
	}
	if target == 0 {
		return true
	}
	if target < 0 {
		return false
	}
	if val, ok := cache[target]; ok {
		return val
	}
	for _, number := range numbers {
		remainder := target - number
		if canSum(remainder, numbers, cache) {
			cache[target] = true
			return cache[target]
		}

	}
	cache[target] = false
	return cache[target]
}

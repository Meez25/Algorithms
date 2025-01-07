package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	target := 100
	listOfAvailableNumbers := []int{4, 3, 5, 25}
	cache := make(map[int][]int, target)
	result := bestSum(target, listOfAvailableNumbers, cache)
	fmt.Println(result, "in", time.Since(start))
}

func bestSum(target int, numbers []int, cache map[int][]int) []int {
	if val, ok := cache[target]; ok {
		return val
	}
	if target == 0 {
		return []int{}
	}
	if target < 0 {
		return nil
	}

	var bestArray []int

	for _, number := range numbers {
		remainder := target - number
		childResult := bestSum(remainder, numbers, cache)
		if childResult != nil {
			childResult = append(childResult, number)
			if len(bestArray) == 0 || len(bestArray) > len(childResult) {
				bestArray = childResult
			}
		}
	}
	if bestArray != nil {
		cache[target] = bestArray
		return bestArray
	}
	cache[target] = nil
	return nil
}

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	target := 300
	listOfAvailableNumbers := []int{7, 14}
	cache := make(map[int]*[]int, target)
	result := howSum(target, listOfAvailableNumbers, cache)
	fmt.Println(result, "in", time.Since(start))
}

func howSum(target int, numbers []int, cache map[int]*[]int) *[]int {
	if val, ok := cache[target]; ok {
		return val
	}
	if target == 0 {
		return &[]int{}
	}
	if target < 0 {
		return nil
	}
	for _, number := range numbers {
		remainder := target - number
		childResult := howSum(remainder, numbers, cache)
		if childResult != nil {
			*childResult = append(*childResult, number)
			cache[target] = childResult
			return childResult
		}
	}
	cache[target] = nil
	return nil
}

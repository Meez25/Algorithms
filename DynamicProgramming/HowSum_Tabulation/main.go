package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	target := 7
	numbers := []int{5, 3, 4}
	result := howSum(target, numbers)
	fmt.Println(result, "in", time.Since(start))
}

func howSum(target int, numbers []int) []int {
	table := make([][]int, target+1)
	for i := range table {
		table[i] = nil
	}

	table[0] = []int{}

	for i := range table {
		if table[i] != nil {
			for _, number := range numbers {
				if i+number < len(table) {
					table[i+number] = table[i]
					table[i+number] = append(table[i+number], number)
				}
			}
		}
	}

	return table[target]
}

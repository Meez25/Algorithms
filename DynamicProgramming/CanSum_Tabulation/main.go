package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	target := 300
	numbers := []int{14, 7}
	result := canSum(target, numbers)
	fmt.Println(result, "in", time.Since(start))
}

func canSum(target int, numbers []int) bool {
	table := make([]bool, target+1)
	table[0] = true

	for i, val := range table {
		if val == true {
			for _, number := range numbers {
				if i+number < len(table) {
					table[i+number] = true
				}
			}
		}
	}

	return table[target]
}

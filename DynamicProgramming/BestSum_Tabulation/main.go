package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	target := 8
	numbers := []int{2, 3, 5}
	result := bestSum(target, numbers)
	fmt.Println(result, "in", time.Since(start))
}

func bestSum(target int, numbers []int) []int {
	table := make([][]int, target+1)
	for i := range table {
		table[i] = nil
	}

	table[0] = []int{}

	for i := range table {
		if table[i] != nil {
			for _, number := range numbers {
				if i+number < len(table) {
					temp := make([]int, 0)
					temp = append(temp, table[i]...)
					temp = append(temp, number)
					if table[i+number] == nil || len(temp) < len(table[i+number]) {
						table[i+number] = temp
					}
				}
			}
		}
	}

	fmt.Println(table)

	return table[target]
}

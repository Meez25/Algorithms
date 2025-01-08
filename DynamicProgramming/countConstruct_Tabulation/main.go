package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	target := "purple"
	words := []string{"purp", "p", "ur", "le", "purpl"}
	result := countConstruct(target, words)
	fmt.Println(result, "in", time.Since(start))
}

func countConstruct(target string, words []string) int {
	table := make([]int, len(target)+1)

	table[0] = 1

	for i := range table {
		if table[i] > 0 {
			for _, word := range words {
				if strings.HasPrefix(target[i:], word) && i+len(word) < len(table) {
					table[i+len(word)] += table[i]
				}
			}
		}
	}

	return table[len(target)]
}

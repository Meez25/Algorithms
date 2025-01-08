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
	result := allConstruct(target, words)
	fmt.Println(result, "in", time.Since(start))
}

func allConstruct(target string, words []string) [][]string {
	table := make([][][]string, len(target)+1)
	table[0] = [][]string{{}}

	for i := 0; i <= len(target); i++ {
		if table[i] != nil {
			for _, word := range words {
				if i+len(word) <= len(target) && strings.HasPrefix(target[i:], word) {
					for _, combination := range table[i] {
						newCombination := append(combination, word)
						table[i+len(word)] = append(table[i+len(word)], newCombination)
					}
				}
			}
		}
	}

	if table[len(target)] == nil {
		return [][]string{}
	}
	return table[len(target)]
}

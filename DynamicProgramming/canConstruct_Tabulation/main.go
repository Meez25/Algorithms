package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	target := "abcdef"
	words := []string{"ab", "abc", "cd", "def", "abcd"}
	result := canConstruct(target, words)
	fmt.Println(result, "in", time.Since(start))
}

func canConstruct(target string, words []string) bool {
	table := make([]bool, len(target)+1)
	for i := range table {
		table[i] = false
	}

	table[0] = true

	for i := range table {
		if table[i] == true {
			for _, word := range words {
				if strings.HasPrefix(target[i:], word) && i+len(word) < len(table) {
					table[i+len(word)] = true
				}
			}
		}
	}

	return table[len(target)]
}

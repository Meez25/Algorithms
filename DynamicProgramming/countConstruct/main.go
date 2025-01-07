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
	cache := make(map[string]int, len(words))
	result := canConstruct(target, words, cache)
	fmt.Println(result, "in", time.Since(start))
}

func canConstruct(target string, words []string, cache map[string]int) int {
	if val, ok := cache[target]; ok {
		return val
	}
	if target == "" {
		return 1
	}

	possibility := 0
	for _, word := range words {
		if strings.HasPrefix(target, word) {
			count := canConstruct(target[len(word):], words, cache)
			possibility += count
		}
	}
	if possibility > 0 {
		cache[target] = possibility
		return possibility
	}

	cache[target] = 0
	return possibility
}

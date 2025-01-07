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
	cache := make(map[string]bool, len(words))
	result := canConstruct(target, words, cache)
	fmt.Println(result, "in", time.Since(start))
}

func canConstruct(target string, words []string, cache map[string]bool) bool {
	if val, ok := cache[target]; ok {
		return val
	}
	if target == "" {
		return true
	}

	for _, word := range words {
		if strings.HasPrefix(target, word) {
			if canConstruct(target[len(word):], words, cache) {
				cache[target] = true
				return true
			}
		}
	}
	cache[target] = false
	return false
}

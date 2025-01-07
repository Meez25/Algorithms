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
	cache := make(map[string][][]string, len(words))
	result := allConstruct(target, words, cache)
	fmt.Println(result, "in", time.Since(start))
}
func allConstruct(target string, wordBank []string, cache map[string][][]string) [][]string {
	if val, ok := cache[target]; ok {
		return val
	}
	if target == "" {
		return [][]string{{}}
	}

	result := make([][]string, 0)

	for _, word := range wordBank {
		if strings.Index(target, word) == 0 {
			suffix := target[len(word):]
			suffixWays := allConstruct(suffix, wordBank, cache)

			targetWays := make([][]string, len(suffixWays))
			for i, way := range suffixWays {
				newWay := make([]string, 0, len(way)+1)
				newWay = append(newWay, word)
				newWay = append(newWay, way...)
				targetWays[i] = newWay
			}
			result = append(result, targetWays...)
		}
	}
	cache[target] = result
	return result
}

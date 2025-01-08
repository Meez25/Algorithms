package main

import "fmt"

func main() {
	graph := map[rune][]rune{
		'f': {'g', 'i'},
		'g': {'h'},
		'h': {},
		'i': {'g', 'k'},
		'j': {'i'},
		'k': {},
	}
	fmt.Println(hasPath(graph, 'f', 'k'))
}

func hasPath(graph map[rune][]rune, src, dst rune) bool {
	if src == dst {
		return true
	}
	for _, val := range graph[src] {
		if hasPath(graph, val, dst) {
			return true
		}
	}
	return false
}

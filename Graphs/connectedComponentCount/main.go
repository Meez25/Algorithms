package main

import "fmt"

func main() {
	visited := make(map[int]bool)
	count := 0
	graph := map[int][]int{
		0: {8, 1, 5},
		1: {0},
		5: {0, 8},
		8: {0, 5},
		2: {3, 4},
		3: {2, 4},
		4: {3, 2},
	}
	for key := range graph {
		if explore(graph, key, visited) {
			count++
		}
	}
	fmt.Println(count)
}

func explore(graph map[int][]int, current int, visited map[int]bool) bool {
	if visited[current] {
		return false
	}

	visited[current] = true

	for _, neighbor := range graph[current] {
		explore(graph, neighbor, visited)
	}
	return true
}

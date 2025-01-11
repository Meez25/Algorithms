package main

import "fmt"

func main() {
	graph := map[int][]int{
		0: {8, 1, 5},
		1: {0},
		5: {0, 8},
		8: {0, 5},
		2: {3, 4},
		3: {2, 4},
		4: {3, 2},
	}
	fmt.Println(largestComponents(graph))
}

func largestComponents(graph map[int][]int) int {
	largest := 0
	visited := make(map[int]bool)
	for key := range graph {
		size := dfs(graph, key, visited)
		if size > largest {
			largest = size
		}
	}
	return largest
}

func dfs(graph map[int][]int, current int, visited map[int]bool) int {
	if visited[current] {
		return 0
	}
	visited[current] = true
	size := 1

	for _, neighbor := range graph[current] {
		size += dfs(graph, neighbor, visited)
	}

	return size
}

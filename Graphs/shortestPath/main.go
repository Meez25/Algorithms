package main

import (
	"fmt"
)

type Node struct {
	value    string
	distance int
}

func main() {
	edges := [][]string{
		{"w", "x"},
		{"x", "y"},
		{"z", "y"},
		{"z", "v"},
		{"w", "v"},
	}
	graph := buildGraph(edges)
	fmt.Println(graph)
	shortestPath := findShortestPath(graph, "w", "z")
	fmt.Println(shortestPath)
}

func buildGraph(edges [][]string) map[string][]string {
	graph := make(map[string][]string)
	for _, pair := range edges {
		if _, ok := graph[pair[0]]; !ok {
			graph[pair[0]] = []string{}
		}
		if _, ok := graph[pair[1]]; !ok {
			graph[pair[1]] = []string{}
		}
		graph[pair[0]] = append(graph[pair[0]], pair[1])
		graph[pair[1]] = append(graph[pair[1]], pair[0])
	}
	return graph
}

func findShortestPath(graph map[string][]string, start, dest string) int {
	queue := make([]Node, 0)
	visited := make(map[string]bool)

	startingNode := Node{value: start, distance: 0}
	queue = append(queue, startingNode)
	visited[startingNode.value] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.value == dest {
			return current.distance
		}
		for _, neighbor := range graph[current.value] {
			nextNode := Node{value: neighbor, distance: current.distance + 1}
			if visited[nextNode.value] {
				continue
			}
			queue = append(queue, nextNode)
			visited[nextNode.value] = true
		}
	}

	return -1
}

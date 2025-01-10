package main

import (
	"fmt"
	"slices"
)

func main() {
	edges := [][]string{
		{"i", "j"},
		{"k", "i"},
		{"m", "k"},
		{"k", "l"},
		{"o", "n"}}

	graph := buildGraph(edges)
	visited := make(map[string]bool, len(graph))
	fmt.Println(undirectedPath(graph, "m", "j", visited))
}

func buildGraph(input [][]string) map[string][]string {
	graph := make(map[string][]string)
	for _, edge := range input {
		a, b := edge[0], edge[1]
		if !slices.Contains(graph[a], b) {
			graph[a] = append(graph[a], b)
		}
		if !slices.Contains(graph[b], a) {
			graph[b] = append(graph[b], a)
		}
	}
	return graph
}

func undirectedPath(graph map[string][]string, src, dst string, visited map[string]bool) bool {
	if _, ok := visited[src]; ok {
		return false
	}
	if src == dst {
		return true
	}
	visited[src] = true
	for _, neighbor := range graph[src] {
		found := undirectedPath(graph, neighbor, dst, visited)
		if found {
			return true
		}
	}
	return false
}

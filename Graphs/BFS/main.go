package main

import "fmt"

func main() {
	graph := map[rune][]rune{
		'a': {'b', 'c'},
		'b': {'d'},
		'c': {'e'},
		'd': {'f'},
		'e': {},
		'f': {},
	}
	BFSPrint(graph, 'a')
}

func BFSPrint(graph map[rune][]rune, source rune) {
	queue := []rune{}

	queue = append(queue, source)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		fmt.Println(string(current))

		for i := range graph[current] {
			queue = append(queue, graph[current][i])
		}
	}
}

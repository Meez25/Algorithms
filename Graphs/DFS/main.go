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
	DFSPrint(graph, 'a')
}

func DFSPrint(graph map[rune][]rune, source rune) {
	stack := []rune{}

	stack = append(stack, source)

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Println(string(current))

		for i := range graph[current] {
			stack = append(stack, graph[current][i])
		}
	}
}

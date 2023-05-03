package graph

import "fmt"

func recursiveCall(adjList map[string][]string, visited map[string]bool, current string, stack []string) []string {
	visited[current] = true

	for _, neighbour := range adjList[current] {
		if _, ok := visited[current]; !ok {
			stack = recursiveCall(adjList, visited, neighbour, stack)
		}
	}

	stack = append(stack, current)
	return stack
}

func DoTopologicalSort() {
	adjList := map[string][]string{
		"0": {},
		"1": {},
		"2": {"3"},
		"3": {"1"},
		"4": {"0", "1"},
		"5": {"0", "2"},
	}

	visited := map[string]bool{}
	stack := []string{}

	for node := range adjList {
		if _, ok := visited[node]; !ok {
			stack = recursiveCall(adjList, visited, node, stack)
		}
	}

	fmt.Print("Topological sort: ")

	for i := len(stack) - 1; i >= 0; i-- {
		fmt.Printf("%s -> ", stack[i])
	}

	fmt.Println()
}

package graph

import "fmt"

func directedCycleCheck(adjList map[string][]string, visited map[string]bool, current string, recursiveStack map[string]bool) bool {
	visited[current] = true
	recursiveStack[current] = true

	for _, neighbour := range adjList[current] {
		if _, ok := recursiveStack[current]; ok {
			return true
		}

		if _, ok := visited[neighbour]; !ok {
			if directedCycleCheck(adjList, visited, neighbour, recursiveStack) {
				return true
			}
		}
	}

	delete(recursiveStack, current)

	return false
}

func DoHasCycleDirected() {
	adjList := map[string][]string{
		"0": {"2"},
		"2": {"3"},
		"3": {"0"},
		"1": {"0"},
	}

	visited := map[string]bool{}
	recursiveStack := map[string]bool{}

	for node := range adjList {
		if directedCycleCheck(adjList, visited, node, recursiveStack) {
			fmt.Printf("Directed Cycle: %s\n", node)
			return
		}
	}

	fmt.Println("Directed Cycle: None")
}

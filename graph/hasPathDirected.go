package graph

import (
	"fmt"
)

func hasPathDfs(source, dest string) bool {
	if source == dest {
		return true
	}

	for _, neighbour := range adjacencyList[source] {
		if hasPathDfs(neighbour, dest) {
			return true
		}
	}

	return false
}

func hasPathBfs(source, dest string) bool {
	queue := []string{source}

	for len(queue) > 0 {
		popped := queue[0]
		queue = queue[1:]

		if popped == dest {
			return true
		}

		queue = append(queue, adjacencyList[popped]...)
	}

	return false
}

func DoHasPathDirected() {
	fmt.Println("Has Path DFS")
	fmt.Printf("b/w a & e: %t\n", hasPathDfs("a", "e"))
	fmt.Printf("b/w a & g: %t\n", hasPathDfs("a", "g"))
	fmt.Println("-----")
	fmt.Println("Has Path BFS")
	fmt.Printf("b/w a & e: %t\n", hasPathBfs("a", "e"))
	fmt.Printf("b/w a & g: %t\n", hasPathBfs("a", "g"))
}

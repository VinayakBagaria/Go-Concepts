package graph

import "fmt"

func hasPath(source, dest string) bool {
	if source == dest {
		return true
	}

	for _, neighbour := range adjacencyList[source] {
		if hasPath(neighbour, dest) {
			return true
		}
	}

	return false
}

func DoHasPath() {
	fmt.Println("Has Path")
	fmt.Printf("b/w a & e: %t\n", hasPath("a", "e"))
	fmt.Printf("b/w a & g: %t\n", hasPath("a", "g"))
}

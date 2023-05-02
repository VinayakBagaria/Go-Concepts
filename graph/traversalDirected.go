package graph

import (
	"fmt"
)

func dfsLoop(source string) {
	stack := []string{source}

	for len(stack) > 0 {
		popped := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Printf("%s -> ", popped)

		stack = append(stack, adjacencyList[popped]...)
	}
}

func dfsRecursive(source string) {
	fmt.Printf("%s -> ", source)

	for _, neighbour := range adjacencyList[source] {
		dfsRecursive(neighbour)
	}
}

func bfsLoop(source string) {
	queue := []string{source}

	for len(queue) > 0 {
		popped := queue[0]
		queue = queue[1:]
		fmt.Printf("%s -> ", popped)

		queue = append(queue, adjacencyList[popped]...)
	}
}

func bfsRecursive(queue []string) {
	if len(queue) == 0 {
		return
	}

	popped := queue[0]
	queue = queue[1:]
	fmt.Printf("%s -> ", popped)

	queue = append(queue, adjacencyList[popped]...)
	
	bfsRecursive(queue)
}

func DoTraversal() {
	fmt.Println("DFS Loop")
	dfsLoop("a")
	fmt.Println("\n-----")
	fmt.Println("DFS Recursive")
	dfsRecursive("a")
	fmt.Println("\n-----")
	fmt.Println("BFS Loop")
	bfsLoop("a")
	fmt.Println("\n-----")
	fmt.Println("BFS Recursive")
	bfsRecursive([]string{"a"})
	fmt.Println()
}

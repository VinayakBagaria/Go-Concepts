package graph

import "fmt"

func dfsLoop(source string) {
	stack := []string{source}

	for len(stack) > 0 {
		popped := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Printf("%s -> ", popped)

		for _, neighbour := range adjacencyList[popped] {
			stack = append(stack, neighbour)
		}
	}
}

func dfsRecursive(source string) {
	fmt.Printf("%s -> ", source)

	for _, neighbour := range adjacencyList[source] {
		dfsRecursive(neighbour)
	}
}

func bfs(source string) {
	queue := []string{source}

	for len(queue) > 0 {
		popped := queue[0]
		queue = queue[1:]
		fmt.Printf("%s -> ", popped)

		for _, neighbour := range adjacencyList[popped] {
			queue = append(queue, neighbour)
		}
	}
}

func DoTraversal() {
	fmt.Println("DFS Loop")
	dfsLoop("a")
	fmt.Println("\nDFS Recursive")
	dfsRecursive("a")
	fmt.Println("\nBFS")
	bfs("a")
	fmt.Println()
}

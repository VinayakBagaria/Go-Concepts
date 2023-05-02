package graph

import "fmt"

const separator = "--------------------------------------"

func DoWork() {
	fmt.Println("Edge array => Adjancency List")
	fmt.Println(MakeAdjacencyList())
	fmt.Println(separator)
	DoTraversal()
	fmt.Println(separator)
	DoHasPathDirected()
	fmt.Println(separator)
	DoHasPathUndirected()
	fmt.Println(separator)
	DoConnectedComponents()
}

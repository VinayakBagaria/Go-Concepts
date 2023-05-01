package graph

import "fmt"

func DoWork() {
	fmt.Println(MakeAdjacencyList())
	fmt.Println("--------------------------------------")
	DoTraversal()
	fmt.Println("--------------------------------------")
	DoHasPathDirected()
	fmt.Println("--------------------------------------")
	DoHasPathUndirected()
}

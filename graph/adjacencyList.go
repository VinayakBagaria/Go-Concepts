package graph

func MakeAdjacencyList(edges [][]string) map[string][]string {
	graph := make(map[string][]string)

	for _, edge := range edges {
		start, end := edge[0], edge[1]

		if _, ok := graph[start]; !ok {
			graph[start] = []string{}
		}
		graph[start] = append(graph[start], end)

		if _, ok := graph[end]; !ok {
			graph[end] = []string{}
		}
		graph[end] = append(graph[end], start)
	}

	return graph
}

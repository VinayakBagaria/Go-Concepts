package graph

import "fmt"

func dfsLand(grid [][]string, r, c int, visited map[string]bool) bool {
	rowInbound := r >= 0 && r < len(grid)
	columnInbound := c >= 0 && c < len(grid[0])

	if !rowInbound || !columnInbound {
		return false
	}

	if grid[r][c] == "W" {
		return false
	}

	position := fmt.Sprintf("%d,%d", r, c)
	if _, ok := visited[position]; ok {
		return false
	}

	visited[position] = true

	// up
	dfsLand(grid, r-1, c, visited)
	// down
	dfsLand(grid, r+1, c, visited)
	// left
	dfsLand(grid, r, c-1, visited)
	// right
	dfsLand(grid, r, c+1, visited)

	return true
}

func DoDistinctIslandCount() {
	grid := [][]string{
		{"W", "L", "W", "W", "W"},
		{"W", "L", "W", "W", "W"},
		{"W", "W", "W", "L", "W"},
		{"W", "W", "L", "L", "W"},
		{"L", "W", "W", "L", "L"},
		{"L", "L", "W", "W", "W"},
	}

	visited := map[string]bool{}
	counter := 0

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid); c++ {
			if dfsLand(grid, r, c, visited) {
				counter += 1
			}
		}
	}

	fmt.Printf("Distinct island count: %d\n", counter)
}

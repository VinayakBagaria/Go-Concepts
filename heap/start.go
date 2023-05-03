package heap

import "fmt"

func insertElement(heap []string, value string) []string {
	heap = append(heap, value)

	current := len(heap) - 1

	for current > 0 {
		parent := current / 2
		if heap[parent] < heap[current] {
			// swap current & parent
			temp := heap[current]
			heap[current] = heap[parent]
			heap[parent] = temp

			current = parent
		} else {
			break
		}
	}

	return heap
}

func DoWork() {
	heap := []string{}
	insertions := []string{"40", "60", "50"}
	for _, node := range insertions {
		heap = insertElement(heap, node)
	}
	fmt.Printf("Max Heap with node order: %v\nHeap made: %v\n", insertions, heap)
}

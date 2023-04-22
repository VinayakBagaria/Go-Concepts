package sorting

/*
For every position,
	Get the smallest element possible in current position and place it here

Best: O(n^2)
Worst: O(n^2)
Memory: O(1)

1. Not stable
2. No memory
*/

func SelectionSort(items []int) []int {
	for i := 0; i < len(items); i++ {
		minIndex := i
		for j := i + 1; j < len(items); j++ {
			if items[j] < items[minIndex] {
				minIndex = j
			}
		}

		temp := items[i]
		items[i] = items[minIndex]
		items[minIndex] = temp
	}

	return items
}

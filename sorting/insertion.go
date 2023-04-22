package sorting

/*
https://blog.boot.dev/golang/insertion-sort-golang

Just like how we arrange playing cards on a hand.

Best: O(n) -> Sorted array: the condition of previous element being bigger than current element fails
							and the inner loop never executes.
Worst: O(n^2)
Memory: O(1)

1. Stable Sort
2. Online -> can sort as a list as it receives
3. No memory extra & In-place
4. No recursion
5. Fast for almost sorted array
*/

func InsertionSort(items []int) []int {
	length := len(items)

	for i := 1; i < length; i++ {
		for j := i; j > 0 && items[j-1] > items[j]; j-- {
			temp := items[j]
			items[j] = items[j-1]
			items[j-1] = temp
		}
	}

	return items
}

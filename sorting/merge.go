package sorting

/*
https://blog.boot.dev/golang/merge-sort-golang
https://github.com/algorithms-go/merge-sort

Divide and Conquer algo

MergeSort() -> recursive
	1. Divide array into 2 equal parts
	2. Call itself on these 2 parts
	3. Merge() call to fit these 2 halves back together

Best: O(n logn)
Worst: O(n logn)
Space: O(n)

1. Extra memory
2. Recursive
3. Stable Sort
*/

func MergeSort(items []int) []int {
	length := len(items)

	if length <= 1 {
		return items
	}

	mid := length / 2
	left := MergeSort(items[:mid])
	right := MergeSort(items[mid:])
	return merge(left, right)
}

func merge(a, b []int) []int {
	i := 0
	j := 0

	final := []int{}

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		final = append(final, a[i])
	}

	for ; j < len(b); j++ {
		final = append(final, b[j])
	}

	return final
}

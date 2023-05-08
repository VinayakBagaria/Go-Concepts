// https://leetcode.com/problems/longest-increasing-subsequence/description/
package random

import "fmt"

type QueueElement struct {
	sequence []int
	index    int
}

func copyFromOneArrayToAnother(source []int) []int {
	var destination []int
	destination = append(destination, source...)
	return destination
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func approach1(nums []int) int {
	queue := []QueueElement{}

	longest := -1

	for index, number := range nums {
		sequence := []int{number}
		queue = append(queue, QueueElement{sequence: sequence, index: index})
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		longest = max(longest, len(node.sequence))

		lastElement := node.sequence[len(node.sequence)-1]
		index := node.index

		for i := index; i < len(nums); i++ {
			if nums[i] > lastElement {
				newSequence := copyFromOneArrayToAnother(node.sequence)
				newSequence = append(newSequence, nums[i])
				queue = append(queue, QueueElement{sequence: newSequence, index: i})
			}
		}
	}

	return longest
}

func maxOfArray(nums []int) int {
	maximum := nums[0]
	for _, num := range nums {
		maximum = max(maximum, num)
	}
	return maximum
}

func lengthOfLIS(nums []int) int {
	// https://www.youtube.com/watch?v=U9_9JmzeaUY
	lis := make([]int, len(nums))
	for i := range lis {
		lis[i] = 1
	}

	for i, num := range nums {
		for j := 0; j < i; j++ {
			if nums[j] < num {
				lis[i] = max(lis[i], lis[j]+1)
			}
		}
	}

	return maxOfArray(lis)
}

func lengthOfLISDp(nums []int) int {
	// https://www.youtube.com/watch?v=odrfUCS9sQk
	dp := make([]int, len(nums))
	dp[0] = 1

	for i, num := range nums {
		maximum := 0
		for j := 0; j < i; j++ {
			if nums[j] < num {
				maximum = max(maximum, dp[j])
			}
		}

		dp[i] = maximum + 1
	}

	return maxOfArray(dp)
}

func DoLongestIncreasingSubsequence() {
	inputs := [][]int{
		{1, 2, 4, 3},
		{10, 9, 2, 5, 3, 7, 101, 18},
	}

	for _, input := range inputs {
		fmt.Printf("Input: %v\n", input)
		fmt.Printf("Naive: %d, Fast: %d; DP: %d\n", approach1(input), lengthOfLIS(input), lengthOfLISDp(input))
	}
}

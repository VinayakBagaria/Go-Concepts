// https://leetcode.com/problems/product-of-array-except-self/
package random

import "fmt"

func productExceptSelf(nums []int) []int {
	left := make([]int, len(nums))
	left[0] = 1

	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	right := make([]int, len(nums))
	right[len(nums)-1] = 1

	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	output := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		output[i] = left[i] * right[i]
	}

	return output
}

func DoProductOfArrayExceptSelf() {
	inputs := [][]int{}
	inputs = append(inputs, []int{1, 2, 3, 4})
	inputs = append(inputs, []int{-1, 1, 0, -3, 3})

	fmt.Println("Product of Array Except Self")

	for _, input := range inputs {
		fmt.Printf("Input: %v, Output: %v\n", input, productExceptSelf(input))
	}
}

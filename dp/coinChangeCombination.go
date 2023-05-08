// https://leetcode.com/problems/coin-change-ii
package dp

import "fmt"

func coinChangeCombination(coins []int, amount int) int {
	// https://www.youtube.com/watch?v=l_nR5X9VmaI
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			dp[j] = dp[j] + dp[j-coin]
		}
	}

	return dp[amount]
}

func DoCoinChangeCombination() {
	case1 := []int{1, 2, 5}
	amount1 := 5

	case2 := []int{2, 3, 5}
	amount2 := 7

	allCases := [][]int{case1, case2}
	allAmounts := []int{amount1, amount2}

	for i, inputCase := range allCases {
		fmt.Printf("Required: %d from coins %v -> %d\n", allAmounts[i], inputCase, coinChangeCombination(inputCase, allAmounts[i]))
	}
}

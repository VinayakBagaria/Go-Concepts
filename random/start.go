package random

import "fmt"

const separator = "--------------------------------------"

func DoWork() {
	fmt.Println(separator)
	DoProductOfArrayExceptSelf()
	fmt.Println(separator)
	DoLongestIncreasingSubsequence()
}

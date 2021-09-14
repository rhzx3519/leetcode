package main

import "fmt"

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = 1
	}
	dp[0] = 1

	for i := 1; i <= n; i++ {
		for j := 1; j < i; j++ {
			if nums[j-1] < nums[i-1] {
				dp[i] = max(dp[i], dp[j] + 1)
			}
		}
	}
	var lis int
	for _, i := range dp {
		lis = max(i, lis)
	}
	return lis
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(lengthOfLIS([]int{10,9,2,5,3,7,101,18}))
}
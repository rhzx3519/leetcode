package main

import (
	"fmt"
	"math"
)

func numSquares(n int) int {
	nums := []int{}
	for i := 1; i*i <= n; i++ {
		nums = append(nums, i*i)
	}

	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}
	for _, num := range nums {
		for i := num; i <= n; i++ {
			dp[i] = min(dp[i], dp[i-num] + 1)
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(numSquares(12))
	fmt.Println(numSquares(13))
}
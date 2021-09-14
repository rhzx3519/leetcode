package main

import (
	"fmt"
	"math"
)

var MOD = int(math.Pow10(9)) + 7

func numWays(steps int, arrLen int) int {
	n := arrLen
	maxC := min(steps, n)
	dp := make([][]int, steps+1)
	for i := 0; i < steps+1; i++ {
		dp[i] = make([]int, maxC+1)
	}
	dp[0][0] = 1

	for i := 1; i <= steps; i++ {
		for j := 0; j <= maxC; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= 1 {
				dp[i][j] = (dp[i][j] + dp[i-1][j-1]) % MOD
			}
			if j < maxC-1 {
				dp[i][j] = (dp[i][j] + dp[i-1][j+1]) % MOD
			}
		}
	}

	return dp[steps][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(numWays(3, 2))
	fmt.Println(numWays(2, 4))
	fmt.Println(numWays(4, 2))
}

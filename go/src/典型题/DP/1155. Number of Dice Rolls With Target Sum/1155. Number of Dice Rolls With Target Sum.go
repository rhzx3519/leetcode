package main

import (
	"fmt"
	"math"
)

/**
	dp[i][j]表示i个骰子，总和为j时的组合总数
 	dp[i][j] = dp[i][j] + dp[i-1][j-k], k为第i个骰子的点数
 */
var mod = int(math.Pow10(9)) + 7

func numRollsToTarget(d int, f int, target int) int {
	dp := make([][]int, 31)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 1001)
	}

	for i := 1; i <= f; i++ {
		dp[1][i] = 1
	}

	for i := 1; i <= d; i++ {
		for j := 1; j <= f; j++ {
			for k := j; k <= target; k++ {
				dp[i][k] = (dp[i][k] + dp[i-1][k - j])%mod
			}
		}
	}
	return dp[d][target]
}

func main() {
	fmt.Println(numRollsToTarget(1,6,3))
	fmt.Println(numRollsToTarget(2,6,7))
	fmt.Println(numRollsToTarget(2,5,10))
	fmt.Println(numRollsToTarget(1,2,3))
	fmt.Println(numRollsToTarget(30,30,500))
}

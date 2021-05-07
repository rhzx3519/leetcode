package main

import (
	"fmt"
	"math"
)

/**
思路：dp[i][k]表示frog跳到点i、跑到k的最小侧跳次数。
dp[i][k] = -1 if obstacles[i]==1
dp[i][k] = min(dp[i-1][k], dp[i-1][other k] + 1)
 */

var INF = math.MaxInt32 - 10

func minSideJumps(obstacles []int) int {
	n := len(obstacles)
	k := 3
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, k)
	}
	dp[0][0] = 1
	dp[0][2] = 1

	for i := 1; i < n; i++ {
		ob := obstacles[i]
		switch ob {
		case 0:
			dp[i][0] = min(dp[i-1][0], min(dp[i-1][1], dp[i-1][2])+1)
			dp[i][1] = min(dp[i-1][1], min(dp[i-1][0], dp[i-1][2])+1)
			dp[i][2] = min(dp[i-1][2], min(dp[i-1][0], dp[i-1][1])+1)
		case 1: // 0
			dp[i][0] = INF
			dp[i][1] = min(dp[i-1][1], dp[i-1][2] + 1)
			dp[i][2] = min(dp[i-1][2], dp[i-1][1] + 1)
		case 2: // 1
			dp[i][1] = INF
			dp[i][0] = min(dp[i-1][0], dp[i-1][2] + 1)
			dp[i][2] = min(dp[i-1][2], dp[i-1][0] + 1)
		case 3: // 2
			dp[i][2] = INF
			dp[i][0] = min(dp[i-1][0], dp[i-1][1] + 1)
			dp[i][1] = min(dp[i-1][1], dp[i-1][0] + 1)
		}

	}
	n--
	return min(dp[n][0], min(dp[n][1], dp[n][2]))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minSideJumps([]int{0,1,2,3,0}))
	fmt.Println(minSideJumps([]int{0,1,1,3,3,0}))
	fmt.Println(minSideJumps([]int{0,2,1,0,3,0}))
	fmt.Println(minSideJumps([]int{0,0,3,1,0,1,0,2,3,1,0}))
}

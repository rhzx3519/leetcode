package main

import "math"

/**
思路：dp[i][j]表示有i+1枚鸡蛋，共有j层楼时，最小操作次数

1. i = 0 即只剩一枚鸡蛋，此时我们需要从 1 层开始逐层验证才能确保获取确切的 f 值，因此对于任意的 j 都有 dp[0][j] = j

2. i = 1，对于任意 j ，第一次操作可以选择在 [1, j] 范围内的任一楼层 k，如果鸡蛋在 k 层丢下后破碎，
接下来问题转化成 i = 0 时验证 k - 1 层需要的次数，即 dp[0][k - 1], 总操作次数为 dp[0][k - 1] + 1；
如果鸡蛋在 k 层丢下后没碎，接下来问题转化成 i = 1 时验证 j - k 层需要的次数， 即 dp[1][j - k],
总操作次数为 dp[1][j - k] + 1，考虑最坏的情况，两者取最大值则有 dp[1][j] = min(dp[1][j], max(dp[0][k - 1] + 1, dp[1][j - k] + 1))
 */

func twoEggDrop(n int) int {
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[0][0] = 0
	dp[1][0] = 0
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}
	for j := 1; j <= n; j++ {
		for k := 1; k <= j; k++ {
			dp[1][j] = min(dp[1][j], max(dp[0][k-1]+1, dp[1][j-k]+1))
		}
	}

	return dp[1][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}









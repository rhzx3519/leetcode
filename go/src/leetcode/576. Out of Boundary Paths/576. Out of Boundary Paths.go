package main

import "math"

/**
https://leetcode-cn.com/problems/out-of-boundary-paths/
思路：dp
dp[k][i][j]表示剩余k步，处于(i,j)位置的路径数
 */

var mod = int(math.Pow10(9)) + 7

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	dp := make([][][]int, maxMove + 1)
	for i := range dp {
		dp[i] = make([][]int, m)
		for j := range dp[i] {
			dp[i][j] = make([]int, n)
		}
	}

	var k1, k2, k3, k4 int
	for k := 1; k <= maxMove; k++ {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if i == 0 {
					k1 = 1
				} else {
					k1 = dp[k-1][i-1][j]
				}

				if i == m-1 {
					k2 = 1
				} else {
					k2 = dp[k-1][i+1][j]
				}

				if j == 0 {
					k3 = 1
				} else {
					k3 = dp[k-1][i][j-1]
				}

				if j == n-1 {
					k4 = 1
				} else {
					k4 = dp[k-1][i][j+1]
				}

				dp[k][i][j] = (k1 + k2 + k3 + k4) % mod
			}
		}
	}

	return dp[maxMove][startRow][startColumn]
}

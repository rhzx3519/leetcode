package main

/*
*   https://leetcode.cn/problems/number-of-ways-of-cutting-a-pizza/description/
 */
func ways(pizza []string, k int) int {
	m := len(pizza)
	n := len(pizza[0])
	mod := 1_000_000_007
	apples := make([][]int, m+1)
	for i := range apples {
		apples[i] = make([]int, n+1)
	}
	dp := make([][][]int, k+1)
	for i := range dp {
		dp[i] = make([][]int, m+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, n+1)
		}
	}

	// 预处理
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			apples[i][j] = apples[i+1][j] + apples[i][j+1] - apples[i+1][j+1]
			if pizza[i][j] == 'A' {
				apples[i][j] += 1
			}
			if apples[i][j] > 0 {
				dp[1][i][j] = 1
			}
		}
	}

	for ki := 2; ki <= k; ki++ {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				// 水平方向切
				for i2 := i + 1; i2 < m; i2++ {
					if apples[i][j] > apples[i2][j] {
						dp[ki][i][j] = (dp[ki][i][j] + dp[ki-1][i2][j]) % mod
					}
				}
				// 垂直方向切
				for j2 := j + 1; j2 < n; j2++ {
					if apples[i][j] > apples[i][j2] {
						dp[ki][i][j] = (dp[ki][i][j] + dp[ki-1][i][j2]) % mod
					}
				}
			}
		}
	}
	return dp[k][0][0]
}

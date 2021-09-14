package main

/**
dp[i][j][0] 表示对长度为[i,j]的石子堆先手取得的最大分数
dp[i][j][1] 表示对长度为[i,j]的石子堆先手取得的最大分数
 */
func stoneGame(piles []int) bool {
	n := len(piles)
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = []int{0, 0}
		}
	}
	for i := 0; i < n; i++ {
		dp[i][i] = []int{piles[i], 0}
	}

	for l := 1; l < n; l++ {	// l 是数组长度
		for i := 0; i < n-l; i++ { // i是数组左端点, j = i+l 是数组右端点
			j := i + l
			left := piles[i] + dp[i+1][j][1]	// 先手取左边的
			right := dp[i][j-1][1] + piles[j]	// 先手取右边的
			if left > right {
				dp[i][j][0] = left
				dp[i][j][1] = dp[i+1][j][0]
			} else {
				dp[i][j][0] = right
				dp[i][j][1] = dp[i][j-1][0]
			}
		}
	}

	return dp[0][n-1][0] > dp[0][n-1][1]
}

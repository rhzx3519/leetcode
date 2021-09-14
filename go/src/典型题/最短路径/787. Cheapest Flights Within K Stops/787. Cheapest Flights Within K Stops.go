package main

import (
	"fmt"
	"math"
)

/**
https://leetcode-cn.com/problems/cheapest-flights-within-k-stops/
f[t][i] 表示经过t次中转，从src到达i城市的最小花费

 */
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	dp := make([][]int, k+2)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32>>1
		}
	}
	dp[0][src] = 0
	for t := 1; t <= k+1; t++ {
		for _, f := range flights {
			u, v, w := f[0], f[1], f[2]
			dp[t][v] = min(dp[t][v], w + dp[t-1][u])
		}
	}

	var ans = math.MaxInt32>>1
	for t := 1; t <= k+1; t++ {
		ans = min(ans, dp[t][dst])
	}
	if ans == math.MaxInt32>>1 {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(findCheapestPrice(3, [][]int{{0,1,100},{1,2,100},{0,2,500}}, 0, 2, 0))
}
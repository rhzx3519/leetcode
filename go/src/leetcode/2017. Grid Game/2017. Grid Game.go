/**
  @author ZhengHao Lou
  Date    2021/09/28
https://leetcode-cn.com/problems/grid-game/
思路：前缀和
*/
package main

import (
	"fmt"
	"math"
)

func gridGame(grid [][]int) int64 {
	m, n := 2, len(grid[0])
	sum := make([][]int64, m)
	for i := range sum {
		sum[i] = make([]int64, n + 1)
		for j := 1; j <= n; j++ {
			sum[i][j] = sum[i][j - 1] + int64(grid[i][j - 1])
		}
	}

	var score int64 = math.MaxInt64 >> 1
	for i := 0; i < n; i++ {
		score = min(score, max(sum[0][n] - sum[0][i+1], sum[1][i]))
	}

	return score
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(gridGame([][]int{{2,5,4},{1,5,1}}))
}


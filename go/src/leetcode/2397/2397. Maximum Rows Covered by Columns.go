/*
*
@author ZhengHao Lou
Date    2022/09/05
*/
package main

import "fmt"

/*
*
https://leetcode.cn/problems/maximum-rows-covered-by-columns/
思路：压缩dp
dp[j]表示当选中状态为j时，被覆盖的最多行数，j的二进制表示被选中的列的状态
*/
func maximumRows(mat [][]int, cols int) int {
	_, n := len(mat), len(mat[0])
	dp := make([]int, 1<<n)
	for i := range mat {
		var mask int
		for j := range mat[i] {
			mask |= mat[i][j] << j
		}
		for j := 0; j < 1<<n; j++ {
			if (j & mask) == mask {
				dp[j]++
			}
		}
	}

	var ans int
	for i := 0; i < 1<<n; i++ {
		var c int
		for j := 0; j < n; j++ {
			if i>>j&1 == 1 {
				c++
			}
		}
		if c == cols {
			ans = max(ans, dp[i])
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maximumRows([][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 0, 1}}, 2))
}

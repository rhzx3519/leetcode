/**
@author ZhengHao Lou
Date    2021/11/22
*/
package main

/**
https://leetcode-cn.com/problems/knight-probability-in-chessboard/
 */

var dxy = [][]int{{2,1},{2,-1},{-2,1},{-2,-1},{1,2},{1,-2},{-1,2},{-1,-2}}
func knightProbability(n int, k int, row int, column int) float64 {
	mem := map[[3]int]float64{}
	var dfs func(r, c, k int) float64
	dfs = func(r, c, k int) float64 {
		if r < 0 || r >= n || c < 0 || c >= n {
			return 0
		}
		if k == 0 {
			return 1
		}
		key := [3]int{r, c, k}
		if p, ok := mem[key]; ok{
			return p
		}
		var p float64
		for i := range dxy {
			p += dfs(r + dxy[i][0], c + dxy[i][1], k - 1)
		}
		p /= 8.0
		mem[key] = p
		return p
	}

	return dfs(row, column, k)
}

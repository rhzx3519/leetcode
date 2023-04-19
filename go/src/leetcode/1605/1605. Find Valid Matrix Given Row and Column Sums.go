package main

/*
*
https://leetcode.cn/problems/find-valid-matrix-given-row-and-column-sums/
*/
func restoreMatrix(rowSum []int, colSum []int) [][]int {
	m, n := len(rowSum), len(colSum)
	g := make([][]int, m)
	for i := range g {
		g[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			g[i][j] = min(rowSum[i], colSum[j])
			rowSum[i] -= g[i][j]
			colSum[j] -= g[i][j]
		}
	}
	return g
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

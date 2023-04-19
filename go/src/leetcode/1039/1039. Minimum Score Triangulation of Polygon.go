package main

import "math"

/*
*
https://leetcode.cn/problems/minimum-score-triangulation-of-polygon/description/
思路：DP
递推公式: f[i][j] = min(v[i]*v[k]*v[j] + f[i][k] + f[k][j]), i < k < j
1. 由于i < k, f[i]要能从f[k]转移过来，必须先计算出f[k],所以i要倒序 i <- k
2. 由于j > k, f[j] 要能从f[k]转移过来，必须先计算出f[k], 所以j要正序 k -> j
*/
func minScoreTriangulation(values []int) int {
	n := len(values)
	f := make([][]int, n)
	for i := range values {
		f[i] = make([]int, n)
	}

	for i := n - 3; i >= 0; i-- {
		for j := i + 2; j < n; j++ {
			f[i][j] = math.MaxInt32
			for k := i + 1; k < j; k++ {
				f[i][j] = min(f[i][j], values[i]*values[k]*values[j]+f[i][k]+f[k][j])
			}
		}
	}
	return f[0][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

package main

import "math"

/*
*
https://leetcode.cn/problems/minimum-difficulty-of-a-job-schedule/
*/
func minDifficulty(jobDifficulty []int, d int) int {
	n := len(jobDifficulty)
	if n < d {
		return -1
	}
	f := make([][]int, d)
	for i := range f {
		f[i] = make([]int, n)
	}
	f[0][0] = jobDifficulty[0]
	for i := 1; i < n; i++ {
		f[0][i] = max(f[0][i-1], jobDifficulty[i])
	}

	for i := 1; i < d; i++ {
		for j := n - 1; j >= i; j-- {
			f[i][j] = math.MaxInt32
			var mx int
			for k := j; k >= i; k-- {
				mx = max(mx, jobDifficulty[k])
				f[i][j] = min(f[i-1][k-1]+mx, f[i][j])
			}
		}
	}
	return f[d-1][n-1]
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

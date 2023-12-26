package main

import (
	"math"
	"math/bits"
)

/*
*
https://leetcode.cn/problems/maximum-students-taking-exam/?envType=daily-question&envId=2023-12-26
*/
func maxStudents(seats [][]byte) int {
	m, n := len(seats), len(seats[0])
	memo := make(map[int]int)

	isSingleRowCompliant := func(status, row int) bool {
		for j := 0; j < n; j++ {
			if (status>>j)&1 == 1 {
				if seats[row][j] == '#' {
					return false
				}
				if j > 0 && (status>>(j-1))&1 == 1 {
					return false
				}
			}
		}
		return true
	}

	isCrossRowCompliant := func(status, upperRowStatus int) bool {
		for j := 0; j < n; j++ {
			if (status>>j)&1 == 1 {
				if j > 0 && (upperRowStatus>>(j-1))&1 == 1 {
					return false
				}
				if j < n-1 && (upperRowStatus>>(j+1))&1 == 1 {
					return false
				}
			}
		}
		return true
	}

	var dp func(int, int) int
	dp = func(row int, status int) int {
		key := (row << n) + status
		if _, ok := memo[key]; !ok {
			if !isSingleRowCompliant(status, row) {
				memo[key] = math.MinInt32
				return memo[key]
			}
			students := bits.OnesCount(uint(status))
			if row == 0 {
				memo[key] = students
				return memo[key]
			}
			var mx int
			for upperRowStatus := 0; upperRowStatus < (1 << n); upperRowStatus++ {
				if isCrossRowCompliant(status, upperRowStatus) {
					mx = max(mx, dp(row-1, upperRowStatus))
				}
			}
			memo[key] = students + mx
		}

		return memo[key]
	}

	var mx int
	for status := 0; status < (1 << n); status++ {
		mx = max(mx, dp(m-1, status))
	}
	return mx
}

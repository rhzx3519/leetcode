package main

import "math"

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	n1, n2 := len(nums1), len(nums2)
	dp := make([][]int, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]int, n2+1)
	}
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			t := 0
			if nums1[i-1] == nums2[j-1] {
				t = 1
			}
			dp[i][j] = max(dp[i-1][j-1] + t, dp[i-1][j], dp[i][j-1])
		}
	}
	return dp[n1][n2]
}

func max(a ...int) int {
	var maxVal = math.MinInt32
	for _, i := range a {
		if i > maxVal {
			maxVal = i
		}
	}
	return maxVal
}

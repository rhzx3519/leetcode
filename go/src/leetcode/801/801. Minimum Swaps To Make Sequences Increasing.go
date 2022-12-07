/**
@author ZhengHao Lou
Date    2022/10/10
*/
package main

/**
https://leetcode.cn/problems/minimum-swaps-to-make-sequences-increasing/
思路：动态规划
f[i][0/1]分别表示不交换/交换nums1[i]和nums2[i]的情况下，使nums1/nums2严格递增的最小交换次数
*/
func minSwap(nums1 []int, nums2 []int) int {
	n := len(nums1)
	f := make([][2]int, n)
	f[0][1] = 1
	for i := 1; i < n; i++ {
		f[i] = [2]int{n, n}
		a1, a2 := nums1[i-1], nums1[i]
		b1, b2 := nums2[i-1], nums2[i]
		if a1 < a2 && b1 < b2 {
			f[i][0] = f[i-1][0]
			f[i][1] = f[i-1][1] + 1
		}
		if a1 < b2 && b1 < a2 {
			f[i][0] = min(f[i][0], f[i-1][1])
			f[i][1] = min(f[i][1], f[i-1][0]+1)
		}
	}

	return min(f[n-1][0], f[n-1][1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

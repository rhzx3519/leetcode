package main

import "math"

/*
*
https://leetcode.cn/problems/o8SXZn/description/
*/
func storeWater(bucket []int, vat []int) int {
	var mx int
	for i := range vat {
		mx = max(mx, vat[i])
	}
	if mx == 0 {
		return 0
	}

	var ans = math.MaxInt32
	for x := 1; x <= mx; x++ { // 枚举蓄水次数
		var y int
		for i := range vat { // 枚举蓄水次数为x的情况下，升级桶次数的总和y
			y += max(0, (vat[i]+x-1)/x-bucket[i])
		}
		ans = min(ans, x+y)
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

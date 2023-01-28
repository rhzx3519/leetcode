/*
*
@author ZhengHao Lou
Date    2023/01/04
*/
package main

/*
*
https://leetcode.cn/problems/maximum-value-at-a-given-index-in-a-bounded-array/
*/
func maxValue(n int, index int, maxSum int) int {
	ans := 1
	rest := maxSum - n
	l, r := index, index
	for l > 0 || r < n-1 {
		length := r - l + 1
		if rest < length {
			break
		}
		rest -= length
		l = max(0, l-1)
		r = min(n-1, r+1)
		ans++
	}
	ans += rest / n
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

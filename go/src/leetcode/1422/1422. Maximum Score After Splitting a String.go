/**
@author ZhengHao Lou
Date    2022/08/14
*/
package main

/**
https://leetcode.cn/problems/maximum-score-after-splitting-a-string/
*/
func maxScore(s string) int {
	var c0, c1 int
	for i := range s {
		if s[i] == '1' {
			c1++
		}
	}
	var ans int
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '1' {
			c1--
		} else {
			c0++
		}
		ans = max(ans, c0+c1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

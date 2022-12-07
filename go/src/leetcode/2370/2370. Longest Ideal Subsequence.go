/**
@author ZhengHao Lou
Date    2022/08/09
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/longest-ideal-subsequence/
思路：f[i]表示以s[i]结尾的最长子序列
*/
func longestIdealString(s string, k int) int {
	var ans int
	f := make([]int, 26)
	for i := range s {
		c := encode(s[i])
		var s int
		for j := max(0, c-k); j < min(c+k+1, 26); j++ {
			s = max(s, f[j])
		}
		f[c] = 1 + s
	}
	for i := range f {
		ans = max(ans, f[i])
	}
	fmt.Println(f, ans)
	return ans
}

func encode(b byte) int {
	return int(b - 'a')
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

func main() {
	longestIdealString("acfgbd", 2)
}

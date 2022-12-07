/**
@author ZhengHao Lou
Date    2022/09/19
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/length-of-the-longest-alphabetical-continuous-substring/
*/
func longestContinuousSubstring(s string) int {
	var ans int
	n := len(s)
	var j int
	for i := 0; i < n; i = j {
		j = i + 1
		for j < n && s[j] == s[j-1]+1 {
			j++
		}
		ans = max(ans, j-i)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestContinuousSubstring("abacaba"))
	fmt.Println(longestContinuousSubstring("abcde"))
}

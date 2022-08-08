/**
@author ZhengHao Lou
Date    2022/01/07
*/
package main

/**
https://leetcode-cn.com/problems/maximum-nesting-depth-of-the-parentheses/
思路：栈
time: O(n)
space: O(1)
 */
func maxDepth(s string) int {
	var l, ans int
	for i := range s {
		switch s[i] {
		case '(':
			l++
		case ')':
			l--
		default:
		}
		ans = max(ans, l)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
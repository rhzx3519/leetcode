/*
*
@author ZhengHao Lou
Date    2022/12/28
*/
package main

/*
*
https://leetcode.cn/problems/minimum-length-of-string-after-deleting-similar-ends/
*/
func minimumLength(s string) int {
	var l, r = 0, len(s) - 1
	for l < r {
		if s[l] != s[r] {
			break
		}
		var x = s[l]
		for l <= r && s[l] == x {
			l++
		}
		for l <= r && s[r] == x {
			r--
		}
	}
	return r - l + 1
}

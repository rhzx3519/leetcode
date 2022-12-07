/*
*
@author ZhengHao Lou
Date    2022/11/28
*/
package main

/*
*
https://leetcode.cn/problems/append-characters-to-string-to-make-subsequence/
*/
func appendCharacters(s string, t string) int {
	var i, j int
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
			j++
		} else {
			i++
		}
	}
	return len(t) - j
}

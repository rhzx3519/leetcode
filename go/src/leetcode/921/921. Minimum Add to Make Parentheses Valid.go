/**
@author ZhengHao Lou
Date    2022/10/04
*/
package main

/**
https://leetcode.cn/problems/minimum-add-to-make-parentheses-valid/
*/
func minAddToMakeValid(s string) int {
	var c int
	var st []byte
	for i := range s {
		switch s[i] {
		case '(':
			st = append(st, s[i])
		case ')':
			if len(st) > 0 && st[len(st)-1] == '(' {
				st = st[:len(st)-1]
			} else {
				c++
			}
		}
	}
	return c + len(st)
}

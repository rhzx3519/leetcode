/**
@author ZhengHao Lou
Date    2022/09/29
*/
package main

import "strings"

/**
https://leetcode.cn/problems/string-rotation-lcci/
*/
func isFlipedString(s1 string, s2 string) bool {
	n1, n2 := len(s1), len(s2)
	if n1 != n2 {
		return false
	}
	return strings.Contains(s2+s2, s1)
}

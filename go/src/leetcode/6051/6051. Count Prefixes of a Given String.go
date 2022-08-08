/**
@author ZhengHao Lou
Date    2022/05/01
*/
package main

import "strings"

/**
https://leetcode-cn.com/problems/count-prefixes-of-a-given-string/
*/
func countPrefixes(words []string, s string) int {
	var c int
	for _, w := range words {
		if strings.HasPrefix(s, w) {
			c++
		}
	}
	return c
}

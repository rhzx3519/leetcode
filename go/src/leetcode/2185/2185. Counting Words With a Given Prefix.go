/**
@author ZhengHao Lou
Date    2022/03/03
*/
package main

import "strings"

/**
https://leetcode-cn.com/problems/counting-words-with-a-given-prefix/
*/
func prefixCount(words []string, pref string) int {
	var c int
	for _, w := range words {
		if strings.HasPrefix(w, pref) {
			c++
		}
	}
	return c
}

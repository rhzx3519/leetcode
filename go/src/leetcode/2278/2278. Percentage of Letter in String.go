/**
@author ZhengHao Lou
Date    2022/05/23
*/
package main

/**
https://leetcode.cn/problems/percentage-of-letter-in-string/
*/
func percentageLetter(s string, letter byte) int {
	n := len(s)
	var c int
	for i := range s {
		if s[i] == letter {
			c++
		}
	}
	return int(float64(c) * 100 / float64(n))
}

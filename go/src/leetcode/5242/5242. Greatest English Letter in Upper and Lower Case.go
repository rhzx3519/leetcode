/**
@author ZhengHao Lou
Date    2022/06/20
*/
package main

import "unicode"

/**
https://leetcode.cn/problems/greatest-english-letter-in-upper-and-lower-case/
*/
const N = 26

func greatestLetter(s string) string {
	letters := make([]int, N)
	for i := range s {
		if unicode.IsLower(rune(s[i])) {
			letters[int(s[i]-'a')] |= 1
		} else {
			letters[int(s[i]-'A')] |= 1 << 1
		}
	}
	for i := N - 1; i >= 0; i-- {
		if letters[i] == 3 {
			return string('A' + i)
		}
	}
	return ""
}

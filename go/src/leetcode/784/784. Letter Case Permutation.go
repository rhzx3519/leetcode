/**
@author ZhengHao Lou
Date    2022/10/30
*/
package main

import (
	"fmt"
	"unicode"
)

/**
https://leetcode.cn/problems/letter-case-permutation/
*/
func letterCasePermutation(s string) (ans []string) {
	n := len(s)
	var dfs func(i int, bytes []byte)
	dfs = func(i int, bytes []byte) {
		if len(bytes) == n {
			ans = append(ans, string(bytes))
			return
		}

		if unicode.IsLetter(rune(s[i])) {
			var r rune
			if unicode.IsLower(rune(s[i])) {
				r = unicode.ToUpper(rune(s[i]))
			} else {
				r = unicode.ToLower(rune(s[i]))
			}
			dfs(i+1, append(bytes, byte(r)))
			dfs(i+1, append(bytes, s[i]))
		} else {
			dfs(i+1, append(bytes, s[i]))
		}
	}

	dfs(0, []byte{})
	return
}

func main() {
	fmt.Println(letterCasePermutation("a1b2"))
}

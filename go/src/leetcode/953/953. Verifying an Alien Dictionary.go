/**
@author ZhengHao Lou
Date    2022/05/17
*/
package main

/**
https://leetcode.cn/problems/verifying-an-alien-dictionary/
*/
func isAlienSorted(words []string, order string) bool {
	ord := make(map[byte]int)
	for i := range order {
		ord[order[i]] = i
	}
	for i := range words {
		if i > 0 {
			for j := 0; j < min(len(words[i-1]), len(words[i])); j++ {
				if ord[words[i-1][j]] > ord[words[i][j]] {
					return false
				} else if ord[words[i-1][j]] < ord[words[i][j]] {
					goto OUT_FOR
				}
			}

			if len(words[i-1]) > len(words[i]) {
				return false
			}
		}
	OUT_FOR:
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

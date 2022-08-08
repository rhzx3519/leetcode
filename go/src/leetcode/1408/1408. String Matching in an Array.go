/**
@author ZhengHao Lou
Date    2022/08/06
*/
package main

import "strings"

/**
https://leetcode.cn/problems/string-matching-in-an-array/
*/
func stringMatching(words []string) []string {
	var ans []string
	for i, w := range words {
		for j := range words {
			if i != j && strings.Contains(words[j], w) {
				ans = append(ans, w)
				break
			}
		}
	}

	return ans
}

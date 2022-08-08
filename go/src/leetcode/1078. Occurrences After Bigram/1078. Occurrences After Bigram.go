/**
@author ZhengHao Lou
Date    2021/12/26
*/
package main

import "strings"

/**
https://leetcode-cn.com/problems/occurrences-after-bigram/
 */
func findOcurrences(text string, first string, second string) []string {
	strs := strings.Split(text, " ")
	n := len(strs)
	var ans = []string{}
	for i := 0; i + 2 < n; i++ {
		if strs[i] == first && strs[i+1] == second {
			ans = append(ans, strs[i+2])
		}
	}
	return ans
}

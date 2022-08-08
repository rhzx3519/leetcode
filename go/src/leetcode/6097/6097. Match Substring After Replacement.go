/**
@author ZhengHao Lou
Date    2022/06/12
*/
package main

/**
https://leetcode.cn/problems/match-substring-after-replacement/
*/
func matchReplacement(s string, sub string, mappings [][]byte) bool {
	mp := ['z' + 1]['z' + 1]bool{}
	for i := range mappings {
		mp[mappings[i][0]][mappings[i][1]] = true
	}
NEXT:
	for i := 0; i <= len(s)-len(sub); i++ {
		for j := 0; j < len(sub); j++ {
			if sub[j] != s[i+j] && !mp[sub[j]][s[i+j]] {
				continue NEXT
			}
		}
		return true
	}
	return false
}

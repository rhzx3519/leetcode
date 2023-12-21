package main

/*
*
https://leetcode.cn/problems/check-if-a-string-is-an-acronym-of-words/?envType=daily-question&envId=2023-12-20
*/
func isAcronym(words []string, s string) bool {
	if len(words) != len(s) {
		return false
	}
	for i := range words {
		if words[i][0] != s[i] {
			return false
		}
	}
	return true
}

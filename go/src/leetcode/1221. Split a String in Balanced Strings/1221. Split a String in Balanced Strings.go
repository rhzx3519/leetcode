package main

/**
https://leetcode-cn.com/problems/split-a-string-in-balanced-strings/

 */
func balancedStringSplit(s string) int {
	var x, count int
	for i := range s {
		if s[i] == 'L' {
			x++
		} else if s[i] == 'R' {
			x--
		}
		if x == 0 {
			count++
		}
	}
	return count
}

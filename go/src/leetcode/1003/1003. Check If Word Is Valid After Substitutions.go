package main

import "fmt"

/*
*
https://leetcode.cn/problems/check-if-word-is-valid-after-substitutions/description/
*/
func isValid(s string) bool {
	var st []byte
	for i := range s {
		if s[i] > 'a' {
			if len(st) == 0 {
				return false
			}
			if s[i]-st[len(st)-1] != 1 {
				return false
			}
		}
		st = append(st, s[i])
		if s[i] == 'c' {
			st = st[:len(st)-3]
		}
	}
	fmt.Println(st)
	return len(st) == 0
}

func main() {
	isValid("aabcbc")
}

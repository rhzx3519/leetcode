package main

import "fmt"

/*
*
https://leetcode.cn/problems/minimum-string-length-after-removing-substrings/?envType=daily-question&envId=2024-01-10
*/
func minLength(s string) int {
	var st []byte
	for i := range s {
		st = append(st, s[i])
		for len(st) >= 2 && (string(st[len(st)-2:]) == "AB" || string(st[len(st)-2:]) == "CD") {
			st = st[:len(st)-2]
		}
	}
	fmt.Println(string(st))
	return len(st)
}

func main() {
	minLength("AABB")
}

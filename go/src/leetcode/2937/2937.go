package main

import "fmt"

/*
*
https://leetcode.cn/problems/make-three-strings-equal/
*/
func findMinimumOperations(s1 string, s2 string, s3 string) int {
	n := min(len(s1), min(len(s2), len(s3)))
	var x int
	for i := 0; i < n; i++ {
		if s1[i] != s2[i] || s2[i] != s3[i] {
			break
		}
		x++
	}
	if x == 0 {
		return -1
	}
	return len(s1) + len(s2) + len(s3) - 3*x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	//fmt.Println(findMinimumOperations("abc", "abb", "ab"))
	fmt.Println(findMinimumOperations("a", "aabc", "a"))
}

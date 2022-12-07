/*
*
@author ZhengHao Lou
Date    2022/11/29
*/
package main

import "fmt"

/*
https://leetcode.cn/problems/minimum-changes-to-make-alternating-binary-string/
*/
func minOperations(s string) int {
	var a, b int
	var x, y = 0, 1
	for i := range s {
		if x != int(s[i]-'0') {
			a++
		}
		if y != int(s[i]-'0') {
			b++
		}

		x, y = y, x
	}
	return min(a, b)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minOperations("0100"))
	fmt.Println(minOperations("10"))
	fmt.Println(minOperations("1111"))
}

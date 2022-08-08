/**
@author ZhengHao Lou
Date    2022/06/11
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/flip-string-to-monotone-increasing/
思路：寻找边界i，使得i之前的1和i之后的0的和最小
time: O(n)
space: O(1)
*/
func minFlipsMonoIncr(s string) int {
	var n0, n1 int
	for i := range s {
		switch s[i] {
		case '0':
			n0++
		case '1':
		}
	}
	var mi = n0 + n1
	for i := range s {
		switch s[i] {
		case '0':
			n0--
		case '1':
			n1++
		}
		if n0+n1 < mi {
			mi = n0 + n1
		}
	}
	return mi
}

func main() {
	//fmt.Println(minFlipsMonoIncr("00110"))
	//fmt.Println(minFlipsMonoIncr("010110"))
	//fmt.Println(minFlipsMonoIncr("00011000"))
	fmt.Println(minFlipsMonoIncr("11011"))
}

/**
@author ZhengHao Lou
Date    2022/04/18
*/
package main

import "strconv"

/**
https://leetcode-cn.com/problems/calculate-digit-sum-of-a-string/
*/
func digitSum(s string, k int) string {
	for len(s) > k {
		var bytes []byte
		for i := 0; i < len(s); i += k {
			var x int
			for j := i; j < min(len(s), i+k); j++ {
				x += int(s[j] - '0')
			}
			bytes = append(bytes, strconv.Itoa(x)...)
		}
		s = string(bytes)
	}
	return s
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	digitSum("11111222223", 3)
}

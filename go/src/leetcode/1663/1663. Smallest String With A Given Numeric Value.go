/**
@author ZhengHao Lou
Date    2023/01/26
*/
package main

import (
	"fmt"
)

/**
https://leetcode.cn/problems/smallest-string-with-a-given-numeric-value/
*/
func getSmallestString(n int, k int) string {
	var bytes []byte
	for i := 1; i <= n; i++ {
		lower := max(1, k-(n-i)*26)
		k -= lower
		bytes = append(bytes, 'a'+byte(lower-1))
	}
	return string(bytes)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

func main() {
	fmt.Println(getSmallestString(3, 27))
	fmt.Println(getSmallestString(5, 73))
}

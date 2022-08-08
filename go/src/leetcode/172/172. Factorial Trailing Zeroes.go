/**
@author ZhengHao Lou
Date    2022/03/25
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/factorial-trailing-zeroes/
*/
func trailingZeroes(n int) int {
	var c int
	for n >= 5 {
		c += n / 5
		n /= 5
	}
	return c
}

func main() {
	fmt.Println(trailingZeroes(6))
}

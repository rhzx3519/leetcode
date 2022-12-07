/**
@author ZhengHao Lou
Date    2022/10/18
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/sum-of-number-and-its-reverse/
*/
func sumOfNumberAndReverse(num int) bool {

	for i := 0; i <= num>>1; i++ {
		if i+reverse(i) == num {
			return true
		}
	}
	return false
}

func reverse(x int) int {
	var c int
	for x != 0 {
		c = c*10 + x%10
		x /= 10
	}
	return c
}

func main() {
	fmt.Println(reverse(140))
}

package main

import (
	"fmt"
)

/**
https://leetcode-cn.com/problems/2-keys-keyboard/
思路：动态规划
类似爬楼梯

ex1. n=5
C	'A'
P	'AA'
P	'AAA'
P	'AAAA'
P	'AAAAA'

ex2. n=9
C	'A'
P	'AA'
P	'AAA'
C	'AAA'		f3
P	'AAAAAA'	f6
P	'AAAAAAAAA'	f9
 */
func minSteps(n int) int {
	f := make([]int, n + 1)
	for i := 2; i <= n; i++ {
		f[i] = i
	}
	f[1] = 0

	for i := 2; i <= n; i++ {
		for j := 2; j * j <= i; j++ {
			if i%j == 0 {
				f[i] = min(f[i], f[j] + f[i/j])
			}
		}
	}
	return f[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minSteps(9))
}
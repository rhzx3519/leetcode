/**
@author ZhengHao Lou
Date    2022/04/02
*/
package main

import "fmt"

/**
https://leetcode-cn.com/contest/biweekly-contest-75/problems/number-of-ways-to-select-buildings/
*/
func numberOfWays(s string) int64 {
	var c int64
	n := len(s)
	left0, right0 := make([]int, n), make([]int, n)
	left1, right1 := make([]int, n), make([]int, n)
	for c0, c1, i := 0, 0, 0; i < n; i++ {
		left0[i] = c0
		left1[i] = c1
		switch s[i] {
		case '0':
			c0++
		case '1':
			c1++
		}
	}

	for c0, c1, i := 0, 0, n-1; i >= 0; i-- {
		right0[i] = c0
		right1[i] = c1
		switch s[i] {
		case '0':
			c0++
		case '1':
			c1++
		}
	}

	for i := 0; i < n; i++ {
		switch s[i] {
		case '0':
			c += int64(left1[i] * right1[i])
		case '1':
			c += int64(left0[i] * right0[i])
		}
	}

	fmt.Println(left0, right0)
	fmt.Println(left1, right1)
	return c
}

func main() {
	numberOfWays("001101")
	numberOfWays("11100")
}

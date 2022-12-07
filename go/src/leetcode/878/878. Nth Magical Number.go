/*
*
@author ZhengHao Lou
Date    2022/11/22
*/
package main

import "fmt"

/*
*
https://leetcode.cn/problems/nth-magical-number/description/
*/
const MOD int = 1e9 + 7

func nthMagicalNumber(n int, a int, b int) int {
	lcm := a / gcd(a, b) * b
	l, r := 0, min(a, b)*n
	for l < r {
		m := l + (r-l)>>1
		if m/a+m/b-m/lcm >= n {
			r = m
		} else {
			l = m + 1
		}
	}
	return l % MOD
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(nthMagicalNumber(1, 2, 3))
	fmt.Println(nthMagicalNumber(4, 2, 3))
}

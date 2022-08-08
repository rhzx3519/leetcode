package main

import "fmt"

/**
https://leetcode-cn.com/problems/ugly-number-iii/
 */
func nthUglyNumber(n int, a int, b int, c int) int {
	ab := lcm(a, b)
	bc := lcm(b, c)
	ac := lcm(a, c)
	abc := lcm(int(ab), c)

	var l, r int64 = int64(min(a, min(b, c))), 2000000009
	var m int64
	for l < r {
		m = l + (r-l)>>1
		if (m/int64(a) + m/int64(b) +m/int64(c) - m/ab - m/bc - m/ac + m/abc) >= int64(n) {
			r = m
		} else {
			l = m + 1
		}
	}
	return int(l)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int64 {
	return int64(a * b / gcd(a, b))
}

func main() {
	fmt.Println(nthUglyNumber(3, 2, 3, 5))
}

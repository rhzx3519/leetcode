/*
*
@author ZhengHao Lou
Date    2022/12/26
*/
package main

import "sort"

/*
*
https://leetcode.cn/problems/minimize-the-maximum-of-two-arrays/
*/
func minimizeSet(d1 int, d2 int, c1 int, c2 int) int {
	lcm := d1 / gcd(d1, d2) * d2
	return sort.Search((c1+c2)*2+1, func(x int) bool {
		l1 := max(c1-(x/d2-x/lcm), 0)
		l2 := max(c2-(x/d1-x/lcm), 0)
		common := x - x/d1 - x/d2 + x/lcm
		return common >= l1+l2
	})
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
@author ZhengHao Lou
Date    2022/06/08
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/valid-boomerang/
*/
func isBoomerang(points [][]int) bool {
	if (points[0][0] == points[1][0] && points[0][1] == points[1][1]) ||
		(points[0][0] == points[2][0] && points[0][1] == points[2][1]) {
		return false
	}
	
	if points[0][0] == points[1][0] && points[0][0] == points[2][0] {
		return false
	}
	
	a1, b1 := points[0][1]-points[1][1], points[0][0]-points[1][0]
	g1 := gcd(a1, b1)
	a1, b1 = a1/g1, b1/g1
	
	a2, b2 := points[0][1]-points[2][1], points[0][0]-points[2][0]
	g2 := gcd(a2, b2)
	a2, b2 = a2/g2, b2/g2
	return fmt.Sprintf("%v/%v", a1, b1) != fmt.Sprintf("%v/%v", a2, b2)
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

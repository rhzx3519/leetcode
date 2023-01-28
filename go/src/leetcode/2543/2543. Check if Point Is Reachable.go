/**
@author ZhengHao Lou
Date    2023/01/22
*/
package main

/**
https://leetcode.cn/problems/check-if-point-is-reachable/description/
*/
func isReachable(targetX int, targetY int) bool {
	g := gcd(targetX, targetY)
	return g&(g-1) == 0
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

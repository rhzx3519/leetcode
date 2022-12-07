/**
@author ZhengHao Lou
Date    2022/09/19
*/
package main

/**
https://leetcode.cn/problems/smallest-even-multiple/
*/
func smallestEvenMultiple(n int) int {
	x := gcd(n, 2)
	if x == 1 {
		return n * 2
	} else if x == 2 {
		return n
	} else if x == n {
		return n
	}
	return x * n
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	smallestEvenMultiple(6)
}

/**
@author ZhengHao Lou
Date    2022/07/29
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/valid-square/
*/

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	if isSquare(p1, p2, p3) {
		return isSquare(p4, p2, p3)
	}
	if isSquare(p1, p2, p4) {
		return isSquare(p3, p2, p4)
	}
	if isSquare(p1, p3, p4) {
		return isSquare(p2, p3, p4)
	}
	return false
}

func isSquare(p1, p2, p3 []int) bool {
	var a, b [2]int
	a[0] = p2[0] - p1[0]
	a[1] = p2[1] - p1[1]

	b[0] = p3[0] - p1[0]
	b[1] = p3[1] - p1[1]

	la, lb := a[0]*a[0]+a[1]*a[1], b[0]*b[0]+b[1]*b[1]
	return a[0]*b[0]+a[1]*b[1] == 0 &&
		la != 0 && la == lb
}

func main() {
	fmt.Println(validSquare([]int{0, 0}, []int{5, 0}, []int{5, 4}, []int{0, 4}))
}

package main

import "fmt"

/**
https://leetcode-cn.com/problems/rotated-digits/

 */
func rotatedDigits(n int) int {
	var count int
	for i := 1; i <= n; i++ {
		if isGood(i) {
			count++
		}
	}
	return count
}

func isGood(x int) bool {
	var t int
	var f bool
	for ; x != 0; x /= 10 {
		t = x % 10
		if t == 2 || t == 5 || t == 6 || t == 9 {
			f = true
		} else if t == 0 || t == 1 || t == 8 {

		} else {
			return false
		}
	}
	return f
}

func main() {
	fmt.Println(rotatedDigits(857))	// 247
	//fmt.Println(rotatedDigits(10))
	//fmt.Println(rotatedDigits(1))
	//fmt.Println(rotatedDigits(2))
}
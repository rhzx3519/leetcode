/**
@author ZhengHao Lou
Date    2022/04/05
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/prime-number-of-set-bits-in-binary-representation/
*/
func countPrimeSetBits(left int, right int) int {
	primes := map[int]bool{
		2:  true,
		3:  true,
		5:  true,
		7:  true,
		11: true,
		13: true,
		17: true,
		19: true,
		23: true,
	}

	var count = func(x int) int {
		var c int
		for ; x != 0; x &= (x - 1) {
			c++
		}
		return c
	}

	var c int
	for i := left; i <= right; i++ {
		if primes[count(i)] {
			c++
		}
	}
	return c
}

func main() {
	fmt.Println(countPrimeSetBits(842, 888))
}

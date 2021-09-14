package main

import "fmt"

/**
https://leetcode-cn.com/problems/concatenation-of-consecutive-binary-numbers/
 */

const MOD int = 1e9 + 7

func concatenatedBinary(n int) int {
	var ans = 0
	var total int
	for i := 1; i <= n; i++ {
		sz := bin(i)

		ans = (ans * (1<<sz) + i ) % MOD
		total += sz
	}
	fmt.Println(ans)
	return ans
}

// Return length of x in binary
func bin(x int) int {
	var sz int
	for x != 0 {
		sz++
		x >>= 1
	}
	return sz
}

func main() {
	//concatenatedBinary(3)
	concatenatedBinary(12)
}

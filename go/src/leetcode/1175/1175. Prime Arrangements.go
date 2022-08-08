/**
@author ZhengHao Lou
Date    2022/06/30
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/prime-arrangements/
*/
const MOD int = 1e9 + 7

func numPrimeArrangements(n int) int {
	var x int
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			x++
		}
	}
	return frac(x) * frac(n-x) % MOD
}

func isPrime(x int) bool {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func frac(x int) int {
	var ans = 1
	for i := 1; i <= x; i++ {
		ans = (ans * i) % MOD
	}
	return ans
}

func main() {
	fmt.Println(numPrimeArrangements(100))
}

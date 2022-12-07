/**
@author ZhengHao Lou
Date    2022/11/15
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/count-ways-to-build-good-strings/
*/
const MOD int = 1e9 + 7

func countGoodStrings(low int, high int, zero int, one int) (tot int) {
	f := make([]int, high+1)
	f[0] = 1
	for i := 1; i <= high; i++ {
		if i >= zero {
			f[i] = (f[i] + f[i-zero]) % MOD
		}
		if i >= one {
			f[i] = (f[i] + f[i-one]) % MOD
		}
	}
	for i := low; i <= high; i++ {
		tot = (tot + f[i]) % MOD
	}
	fmt.Println(tot)
	return
}

func main() {
	countGoodStrings(3, 3, 1, 1)
	countGoodStrings(2, 3, 1, 2)
}

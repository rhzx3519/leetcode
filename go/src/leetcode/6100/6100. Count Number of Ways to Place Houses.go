/**
@author ZhengHao Lou
Date    2022/06/27
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/count-number-of-ways-to-place-houses/
*/
const MOD int = 1e9 + 7

func countHousePlacements(n int) int {
	var f1, f2 = 0, 1
	for i := 0; i < n; i++ {
		f1, f2 = f2, (f1+f2)%MOD
	}
	return (f1 + f2) * (f1 + f2) % MOD
}

func main() {
	fmt.Println(countHousePlacements(2))
}

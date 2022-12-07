/**
@author ZhengHao Lou
Date    2022/11/18
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode.cn/problems/sum-of-subsequence-widths/description/
*/
const MOD int = 1e9 + 7

func sumSubseqWidths(nums []int) (tot int) {
	n := len(nums)
	power2 := make([]int, n+1)
	power2[0] = 1
	for i := 1; i <= n; i++ {
		power2[i] = power2[i-1] * 2 % MOD
	}
	sort.Ints(nums)
	for i, x := range nums {
		tot = (tot + (power2[i]-power2[n-1-i])*x) % MOD
	}
	if tot < 0 {
		tot = (tot + MOD) % MOD
	}
	return tot
}

func main() {
	//fmt.Println(sumSubseqWidths([]int{2, 1, 3}))
	//fmt.Println(sumSubseqWidths([]int{2}))
	fmt.Println(sumSubseqWidths([]int{96, 87, 191, 197, 40, 101, 108, 35, 169, 50, 168, 182, 95, 80, 144, 43, 18, 60, 174, 13, 77, 173, 38, 46, 80, 117, 13, 19, 11, 6, 13, 118, 39, 80, 171, 36, 86, 156, 165, 190, 53, 49, 160, 192, 57, 42, 97, 35, 124, 200, 84, 70, 145, 180, 54, 141, 159, 42, 66, 66, 25, 95, 24, 136, 140, 159, 71, 131, 5, 140, 115, 76, 151, 137, 63, 47, 69, 164, 60, 172, 153, 183, 6, 70, 40, 168, 133, 45, 116, 188, 20, 52, 70, 156, 44, 27, 124, 59, 42, 172}))
}

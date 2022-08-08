/**
@author ZhengHao Lou
Date    2022/05/16
*/
package main

/**
https://leetcode.cn/problems/largest-combination-with-bitwise-and-greater-than-zero/
*/
const N = 31

func largestCombination(candidates []int) (ans int) {
	for i := 0; i < N; i++ {
		var c int
		for _, num := range candidates {
			if num&(1<<i) != 0 {
				c++
			}
		}
		ans = max(ans, c)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

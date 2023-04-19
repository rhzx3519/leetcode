/**
@author ZhengHao Lou
Date    2021/11/12
*/
package main

import (
	"fmt"
	"math"
)

/**
https://leetcode-cn.com/problems/guess-number-higher-or-lower-ii/
思路：记忆化搜索，取所有最坏情况的最小值
 */
const N int = 201
func getMoneyAmount(n int) int {
	var mem = make([][]int, N)
	for i := range mem {
		mem[i] = make([]int, N)
	}

	var dfs func(l, r int) int
	dfs = func(l, r int) int {
		if l >= r {
			return 0
		}
		if mem[l][r] != 0 {
			return mem[l][r]
		}
		var c int = math.MaxInt32 >> 1
		for i := l; i <= r; i++ {
			d := max(dfs(l, i-1), dfs(i+1, r)) + i
			c = min(c, d)
		}
		mem[l][r] = c
		return c
	}

	return dfs(1, n)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(getMoneyAmount(10))
}

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
 */
const N int = 201

// 区间dp
func getMoneyAmount_1(n int) int {
	var f = make([][]int, N)
	for i := range f {
		f[i] = make([]int, N)
	}
	for len := 2; len <= n; len++ {
		for l := 1; l + len - 1 <= n; l++ {
			r := l + len - 1
			var c int
			f[l][r] = 1e9
			for i := l; i <= r; i++ {
				c = max(f[l][i-1], f[i+1][r]) + i
				f[l][r] = min(f[l][r], c)
			}
		}
	}
	return f[1][n]
}

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
	fmt.Println(getMoneyAmount_1(10))
}

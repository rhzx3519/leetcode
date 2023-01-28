/**
@author ZhengHao Lou
Date    2023/01/16
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/increment-submatrices-by-one/
*/
func rangeAddQueries(n int, queries [][]int) [][]int {
	diffs := make([][]int, n)
	g := make([][]int, n)
	for i := range diffs {
		diffs[i] = make([]int, n+1)
		g[i] = make([]int, n)
	}
	for _, q := range queries {
		for r := q[0]; r <= q[2]; r++ {
			diffs[r][q[1]]++
			diffs[r][q[3]+1]--
		}
	}
	for r := 0; r < n; r++ {
		var x int
		for c := 0; c < n; c++ {
			x += diffs[r][c]
			g[r][c] = x
		}
	}
	fmt.Println(g)
	return g
}

func main() {
	rangeAddQueries(3, [][]int{{1, 1, 2, 2}, {0, 0, 1, 1}})
}

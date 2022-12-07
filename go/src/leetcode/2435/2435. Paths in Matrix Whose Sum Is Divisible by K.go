/**
@author ZhengHao Lou
Date    2022/10/10
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/paths-in-matrix-whose-sum-is-divisible-by-k/
*/
const (
	N       = 50
	MOD int = 1e9 + 7
)

func numberOfPaths(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	f := make([][][]int, m+1)
	for i := range f {
		f[i] = make([][]int, n+1)
		for j := range f[i] {
			f[i][j] = make([]int, 50)
		}
	}
	var s = grid[0][0]
	f[0][0][s%k]++

	for i := 1; i < m; i++ {
		s += grid[i][0]
		f[i][0][s%k] = 1
	}
	s = grid[0][0]
	for j := 1; j < n; j++ {
		s += grid[0][j]
		f[0][j][s%k] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			for t := 0; t < N; t++ {
				f[i][j][(grid[i][j]+t)%k] = (f[i][j][(grid[i][j]+t)%k] + f[i-1][j][t] + f[i][j-1][t]) % MOD
			}
		}
	}

	return f[m-1][n-1][0]
}

func main() {
	fmt.Println(numberOfPaths([][]int{{5, 2, 4}, {3, 0, 5}, {0, 7, 2}}, 3))
	fmt.Println(numberOfPaths([][]int{{0, 0}}, 5))
	fmt.Println(numberOfPaths([][]int{{7, 3, 4, 9}, {2, 3, 6, 2}, {2, 3, 7, 0}}, 1))
}

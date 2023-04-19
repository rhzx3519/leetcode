package main

import (
	"fmt"
	"math"
)

/**
https://leetcode-cn.com/problems/minimize-the-difference-between-target-and-chosen-elements/
思路：DP
f[i][t]表示到第i行为止，能否选出和为t
求出选完m行之后，能够得到的所有元素和
 */

var N = 5000
func minimizeTheDifference(mat [][]int, target int) int {
	m, n := len(mat), len(mat[0])
	f := make([][]bool, m)
	for i := range f {
		f[i] = make([]bool, N+1)
	}
	for c := 0; c < n; c++ {
		f[0][mat[0][c]] = true
	}

	for r := 1; r < m; r++ {
		for c := 0; c < n; c++ {
			num := mat[r][c]
			for t := N; t >= num; t-- {
				f[r][t] = f[r][t] || f[r-1][t-num]
			}
		}
	}

	ans := math.MaxInt32
	for t := 0; t <= N; t++ {
		if f[m-1][t] {
			ans = min(ans, abs(t - target))
		}
	}

	return ans
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minimizeTheDifference([][]int{{1,2,3},{4,5,6},{7,8,9}}, 13))
	fmt.Println(minimizeTheDifference([][]int{{1},{2},{3}}, 100))
	fmt.Println(minimizeTheDifference([][]int{{1,2,9,8,7}}, 6))
}
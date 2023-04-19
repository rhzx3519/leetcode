package main

import (
	"fmt"
	"math"
)

/*
*
https://leetcode.cn/problems/minimum-cost-to-merge-stones/description/
思路：区间DP
*/
func mergeStones(stones []int, k int) int {
	n := len(stones)
	if (n-1)%(k-1) != 0 {
		return -1
	}
	f := make([][]int, n)
	sums := make([]int, n+1)
	for i := range stones {
		sums[i+1] = sums[i] + stones[i]
	}

	for i := n - 1; i >= 0; i-- {
		f[i] = make([]int, n)
		for j := i + 1; j < n; j++ {
			f[i][j] = math.MaxInt32
			for m := i; m < j; m += k - 1 {
				f[i][j] = min(f[i][j], f[i][m]+f[m+1][j])
			}
			if (j-i)%(k-1) == 0 {
				f[i][j] += sums[j+1] - sums[i]
			}
		}
	}

	return f[0][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(mergeStones([]int{3, 2, 4, 1}, 2))
}

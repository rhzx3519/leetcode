package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode-cn.com/problems/largest-submatrix-with-rearrangements/
 */
func largestSubmatrix(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	sum := make([][]int, m)
	for i := range sum {
		sum[i] = make([]int, n)
	}

	var ans int
	for j := 0; j < n; j++ {
		var s int
		for i := m - 1; i >= 0; i-- {
			if matrix[i][j] == 1 {
				s++
			} else {
				s = 0
			}
			sum[i][j] = s
		}
	}

	for i := 0; i < m; i++ {
		mapper := make(map[int]int)
		for j := 0; j < n; j++ {
			mapper[sum[i][j]]++
		}
		keys := []int{}
		for k, _ := range mapper {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		var count int
		for i := len(keys) - 1; i >= 0; i-- {
			k := keys[i]
			count += mapper[k]
			ans = max(ans, count * k)
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(largestSubmatrix([][]int{{0,0,1},{1,1,1},{1,0,1}}))
	fmt.Println(largestSubmatrix([][]int{{1,0,1,0,1}}))
	fmt.Println(largestSubmatrix([][]int{{1,1,0},{1,0,1}}))
	fmt.Println(largestSubmatrix([][]int{{0,0},{0,0}}))
}

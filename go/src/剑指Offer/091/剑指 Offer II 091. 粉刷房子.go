/**
@author ZhengHao Lou
Date    2022/06/25
*/
package main

import (
	"fmt"
	"math"
)

/**
https://leetcode.cn/problems/JEj789/
*/
func minCost(costs [][]int) int {
	n := len(costs)
	f := make([][3]int, n+1)
	for i := 1; i <= n; i++ {
		for j := range f[i] {
			f[i][j] = math.MaxInt32
		}
	}
	for i := range costs {
		f[i+1][0] = min(f[i][1], f[i][2]) + costs[i][0]
		f[i+1][1] = min(f[i][0], f[i][2]) + costs[i][1]
		f[i+1][2] = min(f[i][0], f[i][1]) + costs[i][2]
	}
	return min(f[n][0], min(f[n][1], f[n][2]))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// [[17,2,17],[16,16,5],[14,3,19]], 10
	fmt.Println(minCost([][]int{{17, 2, 17}, {16, 16, 5}, {14, 3, 19}}))
	fmt.Println(minCost([][]int{{7, 6, 2}}))
}

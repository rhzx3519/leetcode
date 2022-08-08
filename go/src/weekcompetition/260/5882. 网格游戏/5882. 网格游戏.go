package main

import (
	"fmt"
	"math"
)

/**
https://leetcode-cn.com/contest/weekly-contest-260/problems/grid-game/
time: O(n)
space: O(n)
 */
func gridGame(grid [][]int) int64 {
	n := len(grid[0])
	sum10 := make([]int64, n+1)
	sum11 := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		sum10[i] = sum10[i-1] + int64(grid[0][i-1])
		sum11[i] = sum11[i-1] + int64(grid[1][i-1])
	}

	var d int64 = math.MaxInt64 >> 1
	for j := 0; j < n; j++ {
		// [0][0:j], [1][j:n] 置为0
		//var diff = int64(0)
		//for k := 0; k <= j; k++ {
		//	diff = max(diff, sum11[j] - sum11[k])
		//}
		//for k := j+1; k < n; k++ {
		//	diff = max(diff, sum10[k+1] - sum10[j+1])
		//}

		d = min(d, max(sum11[j] - sum11[0], sum10[n] - sum10[j+1]))
		////fmt.Println(diff)
		//d = min(d, diff)
	}

	fmt.Println(d)
	return d
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func main() {
	gridGame([][]int{{2,5,4},{1,5,1}})
	gridGame([][]int{{1,3,1,15},{1,3,3,1}})
	gridGame([][]int{{20,3,20,17,2,12,15,17,4,15},
		 			 {20,10,13,14,15,5,2,3,14,3}})	// 63
}


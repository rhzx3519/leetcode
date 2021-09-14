package main

import (
	"fmt"
	"math"
)

/**
https://leetcode-cn.com/problems/minimum-cost-to-reach-destination-in-time/
f[t][i] 表示经过时间t，从0节点到达节点i所花的旅行费用
 */

var (
	INF = math.MaxInt32 >> 1
)

func minCost(maxTime int, edges [][]int, passingFees []int) int {
	n := len(passingFees)
	f := make([][]int, maxTime + 1)
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = INF
		}
	}
	f[0][0] = passingFees[0]

	for t := 1; t <= maxTime; t++ {
		for _, edge := range edges {
			x, y, time := edge[0], edge[1], edge[2]
			if t < time {
				continue
			}
			f[t][x] = min(f[t][x], f[t-time][y] + passingFees[x])
			f[t][y] = min(f[t][y], f[t-time][x] + passingFees[y])
		}
	}

	var ans = INF
	for t := 1; t <= maxTime; t++ {
		ans = min(ans, f[t][n-1])
	}
	if ans == INF {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minCost(30, [][]int{{0,1,10},{1,2,10},{2,5,10},{0,3,1},{3,4,10},{4,5,15}},
				[]int{5,1,2,20,20,3}))
}

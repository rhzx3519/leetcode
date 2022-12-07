/**
@author ZhengHao Lou
Date    2022/11/15
*/
package main

import (
	"fmt"
	"math"
)

/**
https://leetcode.cn/problems/most-profitable-path-in-a-tree/
*/
func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	n := len(edges) + 1
	var g = make([][]int, n)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	bobTime := make([]int, n)
	for i := range bobTime {
		bobTime[i] = n
	}
	var dfsBob func(x, fx, t int) bool
	dfsBob = func(x, fx, t int) bool {
		if x == 0 {
			bobTime[x] = t
			return true
		}
		var arrived bool
		for _, nx := range g[x] {
			if nx != fx && dfsBob(nx, x, t+1) {
				arrived = true
			}
		}
		if arrived {
			bobTime[x] = t
		}
		return arrived
	}

	dfsBob(bob, -1, 0)

	var mx = math.MinInt32 >> 1
	g[0] = append(g[0], -1)
	var dfsAlice func(x, fx, t, score int)
	dfsAlice = func(x, fx, t, score int) {
		if t < bobTime[x] {
			score += amount[x]
		} else if t == bobTime[x] {
			score += amount[x] >> 1
		}
		if len(g[x]) == 1 {
			mx = max(mx, score)
			return
		}

		for _, nx := range g[x] {
			if nx != fx {
				dfsAlice(nx, x, t+1, score)
			}
		}
		return
	}

	dfsAlice(0, -1, 0, 0)
	return mx
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//fmt.Println(mostProfitablePath([][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}}, 3, []int{-2, 4, 2, -4, 6}))
	fmt.Println(mostProfitablePath([][]int{{0, 1}}, 1, []int{-2, 4}))
}

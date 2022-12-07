/*
*
@author ZhengHao Lou
Date    2022/12/05
*/
package main

import "fmt"

/*
*
https://leetcode.cn/problems/divide-nodes-into-the-maximum-number-of-groups/
思路：二分图 + bfs
整个图的连通分量可能大于1，判断连通分量内是否是二分图，累加连通分量的深度
*/
func magnificentSets(n int, edges [][]int) (tot int) {
	g := make([][]int, n+1)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}

	time := make([]int, n+1)
	var clock int

	//  求连通分量的最大深度
	var bfs func(i int) int
	bfs = func(idx int) int {
		var que = []int{idx}
		clock++
		time[idx] = clock

		var m int
		for ; len(que) != 0; m++ {
			for tmp := len(que); tmp != 0; tmp-- {
				i := que[0]
				que = que[1:]
				for _, ni := range g[i] {
					if time[ni] != clock {
						time[ni] = clock
						que = append(que, ni)
					}
				}
			}
		}

		return m
	}

	var colors = make([]int, n+1)
	var nodes []int

	// 判断是否是二分图
	var dfs func(idx, color int) bool
	dfs = func(idx, color int) bool {
		if colors[idx] != 0 {
			return colors[idx] == color
		}
		nodes = append(nodes, idx)
		colors[idx] = color
		for _, ni := range g[idx] {
			if !dfs(ni, -color) {
				return false
			}
		}
		return true
	}

	for i := 1; i <= n; i++ {
		if colors[i] != 0 {
			continue
		}

		nodes = nil
		if !dfs(i, 1) {
			return -1
		}
		// 以连通分量内的每一个节点作为起始点做bfs求最大深度
		var mx int
		for _, j := range nodes {
			mx = max(mx, bfs(j))
		}
		tot += mx
	}

	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(magnificentSets(6, [][]int{{1, 2}, {1, 4}, {1, 5}, {2, 6}, {2, 3}, {4, 6}}))
	fmt.Println(magnificentSets(3, [][]int{{1, 2}, {2, 3}, {1, 3}}))
	fmt.Println(magnificentSets(92, [][]int{{67, 29}, {13, 29}, {77, 29}, {36, 29}, {82, 29}, {54, 29}, {57, 29}, {53, 29}, {68, 29}, {26, 29}, {21, 29}, {46, 29}, {41, 29}, {45, 29}, {56, 29}, {88, 29}, {2, 29}, {7, 29}, {5, 29}, {16, 29}, {37, 29}, {50, 29}, {79, 29}, {91, 29}, {48, 29}, {87, 29}, {25, 29}, {80, 29}, {71, 29}, {9, 29}, {78, 29}, {33, 29}, {4, 29}, {44, 29}, {72, 29}, {65, 29}, {61, 29}}))
}

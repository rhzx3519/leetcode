/*
*
@author ZhengHao Lou
Date    2022/12/05
*/
package main

/*
*
https://leetcode.cn/problems/is-graph-bipartite/
*/
func isBipartite(graph [][]int) bool {
	n := len(graph)
	colors := make([]int, n)
	var dfs func(idx, color int) bool
	dfs = func(idx, color int) bool {
		if colors[idx] != 0 {
			return colors[idx] == color
		}
		colors[idx] = color
		for _, ni := range graph[idx] {
			if !dfs(ni, -color) {
				return false
			}
		}
		return true
	}

	for i := 0; i < n; i++ {
		if colors[i] == 0 && !dfs(i, 1) {
			return false
		}
	}

	return true
}

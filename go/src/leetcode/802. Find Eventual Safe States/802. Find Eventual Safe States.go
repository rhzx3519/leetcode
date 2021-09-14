package main

import "fmt"

/**
https://leetcode-cn.com/problems/find-eventual-safe-states/
思路：dfs, 剪枝

 */
func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	circle := make([]int, n) // (-1, 0, 1) 分别表示在环中、未访问、不在环中

	var dfs func(idx int) bool
	dfs = func(idx int) bool {
		if circle[idx] == -1 {
			return false
		}
		if circle[idx] == 1 {
			return true
		}
		circle[idx] = -1
		for _, ni := range graph[idx] {
			if !dfs(ni) {
				return false
			}
		}
		circle[idx] = 1
		return true
	}

	for i := 0; i < n; i++ {
		if circle[i] == 0 {
			dfs(i)
		}
	}

	var ans = []int{}

	for i := 0; i < n; i++ {
		if circle[i] == 1 {
			ans = append(ans, i)
		}
	}

	fmt.Println(ans)
	return ans
}

func main() {
	eventualSafeNodes([][]int{{1,2},{2,3},{5},{0},{5},{},{}})
	eventualSafeNodes([][]int{{1,2,3,4},{1,2},{3,4},{0,4},{}})
	eventualSafeNodes([][]int{{0},{2,3,4},{3,4},{0,4},{}})
}



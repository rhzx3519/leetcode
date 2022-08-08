package main

import (
	"fmt"
	"math"
)

/**
https://leetcode-cn.com/problems/shortest-path-visiting-all-nodes/

思路： BFS + 状态压缩
 */
func shortestPathLength(graph [][]int) int {
	n := len(graph)
	dist := make([][]int, 1<<n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
		}
	}

	que := [][]int{}
	for i := 0; i < n; i++ {
		dist[1<<i][i] = 0
		que = append(que, []int{1<<i, i})
	}

	for len(que) > 0 {
		t := que[0]
		que = que[1:]
		state, i := t[0], t[1]
		if state == 1<<n - 1 {
			return dist[state][i]
		}

		for _, ni := range graph[i] {
			if dist[state|1<<ni][ni] != math.MaxInt32 {
				continue
			}
			dist[state|1<<ni][ni] = dist[state][i] + 1
			que = append(que, []int{state|1<<ni, ni})
		}
	}

	return -1
}

func main() {
	fmt.Println(shortestPathLength([][]int{{1,2,3},{0},{0},{0}}))
	fmt.Println(shortestPathLength([][]int{{1},{0,2,4},{1,3,4},{2},{1,2}}))
}
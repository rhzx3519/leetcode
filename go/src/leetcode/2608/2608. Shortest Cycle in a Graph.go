package main

import (
	"fmt"
)

/*
*
https://leetcode.cn/problems/shortest-cycle-in-a-graph/
*/
func findShortestCycle(n int, edges [][]int) int {
	var tot = -1
	g := make([][]int, n)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}

	var bfs func(start int)
	bfs = func(start int) {
		dis := make([]int, n)
		for i := 0; i < n; i++ {
			dis[i] = -1
		}
		dis[start] = 0
		type pair struct{ x, fa int }
		que := []pair{{start, -1}}
		for len(que) != 0 {
			x, fa := que[0].x, que[0].fa
			que = que[1:]
			for _, nx := range g[x] {
				if dis[nx] < 0 {
					dis[nx] = dis[x] + 1
					que = append(que, pair{nx, x})
				} else if nx != fa {
					if tot == -1 || dis[x]+dis[nx]+1 < tot {
						tot = dis[x] + dis[nx] + 1
					}
				}
			}
		}
	}

	for start := 0; start < n; start++ {
		bfs(start)
	}

	return tot
}

func main() {
	fmt.Println(findShortestCycle(7, [][]int{{0, 1}, {1, 2}, {2, 0}, {3, 4}, {4, 5}, {5, 6}, {6, 3}}))
}

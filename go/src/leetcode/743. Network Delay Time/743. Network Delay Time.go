package main

import (
	"fmt"
	"math"
)

type edge struct {
	v, w int
}

func networkDelayTime(times [][]int, n int, k int) int {
	adj := make([][]edge, n+1)
	for _, time := range times {
		u, v, w := time[0], time[1], time[2]
		adj[u] = append(adj[u], edge{v: v, w: w})
	}

	var d int
	dis := dijkstra(adj, n, k)
	for i := 1; i <= n; i++ {
		if dis[i] == math.MaxInt32 {
			return -1
		}
		d = max(d, dis[i])
	}
	return d
}

// dijkstra最短路径算法
func dijkstra(adj [][]edge, n, k int) []int {
	vis := make([]bool, n+1)
	dis := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dis[i] = math.MaxInt32
	}

	for _, e := range adj[k] {
		dis[e.v] = e.w
	}

	vis[k] = true
	for i := 1; i <= n; i++ {
		minIdx := -1
		minDis := math.MaxInt32
		for j := 1; j <= n; j++ {
			if vis[j] {
				continue
			}
			if dis[j] < minDis {
				minDis = dis[j]
				minIdx = j
			}
		}
		if minIdx == -1 {
			continue
		}
		vis[minIdx] = true
		for _, e := range adj[minIdx]	{
			dis[e.v] = min(dis[e.v], dis[minIdx] + e.w)
		}
	}
	dis[k] = 0
	return dis
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(networkDelayTime([][]int{{2,1,1},{2,3,1},{3,4,1}}, 4, 2))
	fmt.Println(networkDelayTime([][]int{{1,2,1}}, 2, 1))
	fmt.Println(networkDelayTime([][]int{{1,2,1}}, 2, 2))
	fmt.Println(networkDelayTime([][]int{{1,2,1},{2,1,3}}, 2, 2))
	fmt.Println(networkDelayTime([][]int{{1,2,1},{2,3,2},{1,3,4}}, 3, 1))	// 3
}
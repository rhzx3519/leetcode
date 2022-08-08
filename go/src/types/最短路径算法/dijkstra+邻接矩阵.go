/**
@author ZhengHao Lou
Date    2021/12/03
*/
package 最短路径算法

import "math"

/**
不使用优先队列，适用于稠密图：
1. 不使用优先队列，但引入visited向量记录是否使用过了（使用过的意思是被当做中间节点 ）
2. 使用邻接矩阵表示邻接关系
*/
const INF = math.MaxInt32 >> 1

func dijkstra_(grid [][]int, n, k int) []int {
	vis := make([]bool, n+1)
	dis := make([]int, n+1)

	for v, w := range grid[k] {
		dis[v] = w
	}

	vis[k] = true
	for i := 1; i <= n; i++ {
		minIdx := -1
		minDis := INF
		for j := 1; j <= n; j++ {
			if !vis[j] && dis[j] < minDis {
				minDis = dis[j]
				minIdx = j
			}
		}
		if minIdx == -1 {
			continue
		}
		vis[minIdx] = true
		for v, w := range grid[minIdx] {
			dis[v] = min(dis[v], dis[minIdx]+w)
		}
	}
	return dis
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

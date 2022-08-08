/**
@author ZhengHao Lou
Date    2021/12/03
*/
package 最短路径算法

import "math"

/**
使用优先队列，适用于稀疏图：
想象将所有点分成visited和unvisited两个集合，优先队列记录的是和visited的点相连的边（注意纪录的是边）,队头是这些边中权值最小的一条边。
 */
func dijkstra(x, n int, adj map[int][][]int) []int {
	dis := make([]int, n)
	for i := range dis {
		dis[i] = math.MaxInt32 >> 1
	}
	dis[x] = 0
	que := NewPriorityQueue(func(x, y T) int {
		a, b := x.([]int), y.([]int)
		if a[1] < b[1] {
			return 1
		} else if a[1] > b[1] {
			return -1
		} else {
			return 0
		}
	})
	que.Offer([]int{x, 0})
	for !que.IsEmpty() {
		p := que.Poll().([]int)
		i, cost := p[0], p[1]

		if dis[i] < cost {
			continue
		}
		for _, e := range adj[i] {
			ni, cost := e[0], e[1]
			if dis[ni] > dis[i] + cost {
				dis[ni] = dis[i] + cost
				que.Offer([]int{ni, dis[ni]})
			}
		}
	}

	return dis
}


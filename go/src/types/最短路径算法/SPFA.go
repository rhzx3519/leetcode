/**
@author ZhengHao Lou
Date    2021/12/03
*/
package 最短路径算法

/**
 SPFA
这种算法借助一个队列，尝试通过队首的点，对所有的点进行relax
*/

func SPFA(grid [][]int, n, k int) []int {
	dist := make([]int, n)
	for i := range dist {
		dist[i] = INF
	}
	dist[k] = 0
	que := NewPriorityQueue(Reversed(IntComparator))
	que.Offer(k)
	for !que.IsEmpty() {
		x := que.Poll().(int)
		for i := 0; i < n; i++ { // 尝试对所有的点进行relax
			if dist[x]+grid[x][i] < dist[i] {
				dist[i] = dist[x] + grid[x][i]
				que.Offer(i)
			}
		}
	}
	return dist
}

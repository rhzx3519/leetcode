/**
@author ZhengHao Lou
Date    2022/01/24
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/second-minimum-time-to-reach-destination/
思路：BFS
使用dis[i][0], dis[i][1]分别记录从1到i的最短路径和次短路径长度,
入队的判断条件是当前长度d存在队列中，如果d < dis[i][0], 则为最短路径；
如果dis[i][0] < d && d < dis[i][1], 则为次短路径
*/
const INF int = 1e9 + 7

func secondMinimum(n int, edges [][]int, time int, change int) int {
	adj := make([][]int, n+1)
	for i := range adj {
		adj[i] = []int{}
	}
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}
	dis := make([][]int, n+1)
	for i := range dis {
		dis[i] = []int{INF, INF}
	}
	dis[1] = []int{0, INF}
	que := [][]int{{1, 0}}
	for dis[n][1] == INF {
		p := que[0]
		que = que[1:]
		i, d := p[0], p[1]

		var isRed = (d/change)&1 == 1
		if isRed {
			d += (change - d%change)
		}
		d += time

		for _, ni := range adj[i] {
			if d < dis[ni][0] {
				dis[ni][0] = d
				que = append(que, []int{ni, d})
			} else if dis[ni][0] < d && d < dis[ni][1] {
				dis[ni][1] = d
				que = append(que, []int{ni, d})
			}
		}
	}

	fmt.Println(dis)
	return dis[n][1]
}

func main() {
	secondMinimum(5, [][]int{{1, 2}, {1, 3}, {1, 4}, {3, 4}, {4, 5}}, 3, 5)
	secondMinimum(2, [][]int{{1, 2}}, 3, 2)
}

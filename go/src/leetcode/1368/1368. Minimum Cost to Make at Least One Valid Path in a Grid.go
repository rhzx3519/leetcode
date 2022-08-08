/**
@author ZhengHao Lou
Date    2022/05/31
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/minimum-cost-to-make-at-least-one-valid-path-in-a-grid/
*/
const INF int = 1e9 + 7

var dirs = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func minCost(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dis := make([]int, m*n)
	for i := range dis {
		dis[i] = INF
	}
	dis[0] = 0
	var que = []int{0}
	for len(que) != 0 {
		t := que[0]
		que = que[1:]
		x, y := t/n, t%n
		if x == m-1 && y == n-1 {
			return dis[t]
		}
		for i, dir := range dirs {
			nx, ny := x+dir[0], y+dir[1]
			if nx < 0 || nx == m || ny < 0 || ny == n {
				continue
			}
			d := bool2int(i+1 != grid[nx][ny])
			if dis[t]+d < dis[nx*n+ny] {
				dis[nx*n+ny] = dis[t] + d
				if d == 0 {
					que = append(que, 0)
					copy(que[1:], que[0:])
					que[0] = nx*n + ny
				} else {
					que = append(que, nx*n+ny)
				}
			}
		}
	}

	fmt.Println(dis[m*n-1])
	return dis[m*n-1]
}

func bool2int(f bool) int {
	if f {
		return 1
	}
	return 0
}

func main() {
	//minCost([][]int{{1, 1, 1, 1}, {2, 2, 2, 2}, {1, 1, 1, 1}, {2, 2, 2, 2}})
	minCost([][]int{{1, 1, 3}, {3, 2, 2}, {1, 1, 4}})
}

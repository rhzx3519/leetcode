package main

import (
	"container/heap"
	"fmt"
	"math"
)

/*
*
https://leetcode.cn/problems/minimum-time-to-visit-a-cell-in-a-grid/
*/
var dirs = []struct{ x, y int }{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func minimumTime(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	if grid[0][1] > 1 && grid[1][0] > 1 {
		return -1
	}
	dis := make([][]int, m)
	for i := range dis {
		dis[i] = make([]int, n)
		for j := range dis[i] {
			dis[i][j] = math.MaxInt32
		}
	}
	dis[0][0] = 0

	h := hp{{}}
	for {
		tp := heap.Pop(&h).(tuple)
		d, i, j := tp.d, tp.i, tp.j
		if i == m-1 && j == n-1 {
			return d
		}
		for _, q := range dirs {
			x, y := i+q.x, j+q.y
			if x >= 0 && x < m && y >= 0 && y < n {
				nd := max(d+1, grid[x][y])
				nd += (nd - x - y) % 2 // nd 必须与x+y的奇偶性相同
				if nd < dis[x][y] {
					dis[x][y] = nd // 更新最短路
					heap.Push(&h, tuple{nd, x, y})
				}
			}
		}
	}
}

type tuple struct{ d, i, j int }
type hp []tuple

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].d < h[j].d }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minimumTime([][]int{{0, 1, 3, 2}, {5, 1, 2, 5}, {4, 3, 8, 6}}))
}

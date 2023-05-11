package main

import (
	"container/heap"
	"fmt"
	"math"
)

/*
*
https://leetcode.cn/problems/minimum-moves-to-move-a-box-to-their-target-location/description/
思路：BFS
*/
var diff = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func minPushBox(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	start := tuple{}
	for i := range grid {
		for j := range grid[i] {
			switch grid[i][j] {
			case 'S':
				start.i, start.j = i, j
				grid[i][j] = '.'
			case 'B':
				start.x, start.y = i, j
				grid[i][j] = '.'
			}
		}
	}
	math.MaxUint64
	visit := make(map[tuple]bool)
	visit[start] = true
	que := hp{start}
	for que.Len() != 0 {
		t := heap.Pop(&que).(tuple)

		for _, tmp := range diff {
			ni, nj := t.i+tmp[0], t.j+tmp[1]
			nx, ny := t.x, t.y
			d := t.val
			if ni < 0 || nj < 0 || ni >= m || nj >= n || grid[ni][nj] == '#' {
				continue
			}
			if ni == nx && nj == ny {
				nx, ny = nx+tmp[0], ny+tmp[1]
				if nx < 0 || ny < 0 || nx >= m || ny >= n || grid[nx][ny] == '#' {
					continue
				}
				d++
			}
			if grid[nx][ny] == 'T' {
				return d
			}
			tp := tuple{0, ni, nj, nx, ny}
			if visit[tp] {
				continue
			}
			visit[tp] = true
			heap.Push(&que, tuple{d, ni, nj, nx, ny})
		}
	}

	return -1
}

type tuple struct{ val, i, j, x, y int }
type hp []tuple

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].val < h[j].val }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func main() {
	fmt.Println(minPushBox([][]byte{{'#', '#', '#', '#', '#', '#'},
		{'#', 'T', '#', '#', '#', '#'},
		{'#', '.', '.', 'B', '.', '#'},
		{'#', '.', '#', '#', '.', '#'},
		{'#', '.', '.', '.', 'S', '#'},
		{'#', '#', '#', '#', '#', '#'}}))
}

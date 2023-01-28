/*
*
@author ZhengHao Lou
Date    2022/12/14
*/
package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/*
*
https://leetcode.cn/problems/checking-existence-of-edge-length-limited-paths/
思路：离线查找，并查集

### related questions:
- [2503. 矩阵查询可获得的最大分数](https://leetcode.cn/problems/maximum-number-of-points-from-grid-queries/)
*/
func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	// 并查集
	var parent = make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if px := parent[x]; px != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	var merge func(a, b int)
	merge = func(a, b int) {
		pa, pb := find(a), find(b)
		if pa != pb {
			parent[pb] = pa
		}
	}

	// edges进去小顶堆
	que := hp{}
	for _, e := range edgeList {
		heap.Push(&que, tuple{e[2], e[0], e[1]})
	}
	fmt.Println(que)

	// queries根据limit排序
	var ans = make([]bool, len(queries))
	var id = make([]int, len(queries))
	for i := range id {
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool {
		return queries[id[i]][2] < queries[id[j]][2]
	})
	for _, i := range id {
		p, q, limit := queries[i][0], queries[i][1], queries[i][2]
		for len(que) != 0 && que[0].val < limit {
			tp := heap.Pop(&que).(tuple)
			merge(tp.i, tp.j)
		}
		//fmt.Println(parent)
		ans[i] = find(p) == find(q)
	}

	fmt.Println(ans)
	return ans
}

type tuple struct{ val, i, j int }
type hp []tuple

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].val < h[j].val }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func main() {
	distanceLimitedPathsExist(3, [][]int{{0, 1, 2}, {1, 2, 4}, {2, 0, 8}, {1, 0, 16}}, [][]int{{0, 1, 2}, {0, 2, 5}})
	distanceLimitedPathsExist(5, [][]int{{0, 1, 10}, {1, 2, 5}, {2, 3, 9}, {3, 4, 13}}, [][]int{{0, 4, 14}, {1, 4, 13}})
}

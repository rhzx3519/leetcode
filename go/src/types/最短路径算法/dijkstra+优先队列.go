/*
*
@author ZhengHao Lou
Date    2021/12/03
*/
package 最短路径算法

import (
	"container/heap"
	"math"
)

type pair struct {
	x, cost int
}

func dijkstra1(x, n int, adj map[int][][]int) []int {
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
			if dis[ni] > dis[i]+cost {
				dis[ni] = dis[i] + cost
				que.Offer([]int{ni, dis[ni]})
			}
		}
	}

	return dis
}

func dijkstra(from, to, n int, adj [][]pair) []int {
	dis := make([]int, n)
	for i := range dis {
		dis[i] = math.MaxInt32 >> 1
	}
	dis[from] = 0

	h := hp{{from, 0}}
	heap.Init(&h)

	for h.Len() > 0 {
		p := heap.Pop(&h).(pair)

		if dis[p.x] < p.cost {
			continue
		}
		for _, e := range adj[p.x] {
			ni, cost := e.x, e.cost
			if dis[ni] > dis[p.x]+cost {
				dis[ni] = dis[p.x] + cost
				heap.Push(&h, pair{ni, dis[ni]})
			}
		}
	}

	return dis
}

type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].cost < h[j].cost }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

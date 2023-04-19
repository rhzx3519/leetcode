package main

import (
	"container/heap"
	"math"
)

/*
*
https://leetcode.cn/problems/design-graph-with-shortest-path-calculator/
*/
type pair struct {
	x, cost int
}
type Graph struct {
	n   int
	adj [][]pair
}

func Constructor(n int, edges [][]int) Graph {
	adj := make([][]pair, n)
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], pair{edge[1], edge[2]})
	}
	return Graph{
		n:   n,
		adj: adj,
	}
}

func (this *Graph) AddEdge(edge []int) {
	this.adj[edge[0]] = append(this.adj[edge[0]], pair{edge[1], edge[2]})
}

func (this *Graph) ShortestPath(node1 int, node2 int) int {
	dis := dijkstra(node1, node2, this.n, this.adj)
	if dis[node2] == math.MaxInt32>>1 {
		return -1
	}
	return dis[node2]
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

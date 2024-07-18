package main

import "container/heap"

/**
https://leetcode.cn/problems/minimum-time-to-visit-disappearing-nodes/?envType=daily-question&envId=2024-07-18
*/
func minimumTime(n int, edges [][]int, disappear []int) []int {
    type edge struct{ to, wt int }
    g := make([][]edge, n) // 稀疏图用邻接表
    for _, e := range edges {
        x, y, wt := e[0], e[1], e[2]
        g[x] = append(g[x], edge{y, wt})
        g[y] = append(g[y], edge{x, wt})
    }

    dis := make([]int, n)
    for i := range dis {
        dis[i] = -1
    }
    dis[0] = 0
    h := hp{{}}
    for len(h) > 0 {
        p := heap.Pop(&h).(pair)
        dx := p.dis
        x := p.x
        if dx > dis[x] { // x 之前出堆过
            continue
        }
        for _, e := range g[x] {
            y := e.to
            newDis := dx + e.wt
            if newDis < disappear[y] && (dis[y] < 0 || newDis < dis[y]) {
                dis[y] = newDis // 更新 x 的邻居的最短路
                heap.Push(&h, pair{newDis, y})
            }
        }
    }
    return dis
}

type pair struct{ dis, x int }
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

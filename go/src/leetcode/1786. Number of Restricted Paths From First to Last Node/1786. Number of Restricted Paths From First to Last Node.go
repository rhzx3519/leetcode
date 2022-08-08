package main

import (
	"fmt"
	"math"
)

const (
	INF = math.MaxInt32 >> 1
	MOD int = 1e9 + 7
)

func countRestrictedPaths(n int, edges [][]int) int {
	var adj = buildAdj(n, edges)
	dist := dijkstra(n-1, n, adj)

	mem := make([]int, n)
	for i := range mem {
		mem[i] = -1
	}

	var dp func(idx int) int
	dp = func(idx int) int {
		if mem[idx] != -1 {
			return mem[idx]
		}
		if idx == n - 1 {
			return 1
		}

		var count int
		for _, edge := range adj[idx] {
			v, _ := edge[0], edge[1]
			if dist[v] < dist[idx] {
				count += dp(v)
				count %= MOD
			}
		}

		mem[idx] = count
		return count
	}

	var path = dp(0)
	fmt.Println(dist, path)

	return path
}

func buildAdj(n int, edges [][]int) map[int][][]int {
	adj := make(map[int][][]int)
	for i := 0; i < n; i++ {
		adj[i] = [][]int{}
	}
	for _, edge := range edges {
		u, v, w := edge[0] - 1, edge[1] - 1, edge[2]
		adj[u] = append(adj[u], []int{v, w})
		adj[v] = append(adj[v], []int{u, w})
	}
	return adj
}

func dijkstra(x, n int, adj map[int][][]int) []int {
	dis := make([]int, n)
	for i := range dis {
		dis[i] = INF
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
			if dist[x] + grid[x][i] < dist[i] {
				dist[i] = dist[x] + grid[x][i]
				que.Offer(i)
			}
		}
	}
	return dist
}

type (
	T interface {}
	Comparator func(x, y T) int
)

type Sortable struct {
	List []T
	Cmp  Comparator
}

var (
	// 大顶堆
	IntComparator = func(x, y T) int {
		a, b := x.(int), y.(int)
		if a < b {
			return -1
		} else if a > b {
			return 1
		} else {
			return 0
		}
	}

	// return a reversed version of current Comparator function
	Reversed = func(cmp Comparator) Comparator {
		return func(x, y T) int {
			return -cmp(x, y)
		}
	}
)


// Len is the number of elements in the collection.
func (a *Sortable) Len() int {
	return len(a.List)
}
// Less reports whether the element with
// index i should sort before the element with index j.
func (a *Sortable) Less(i, j int) bool {
	return a.Cmp(a.List[i], a.List[j]) < 0
}

// Swap swaps the elements with indexes i and j.
func (a *Sortable) Swap(i, j int) {
	a.List[i], a.List[j] = a.List[j], a.List[i]
}


type PriorityQueue struct {
	a *Sortable
}

func NewPriorityQueue(comparator Comparator) *PriorityQueue {
	que := &PriorityQueue{
		a: &Sortable{
			List: make([]T, 1),
			Cmp: comparator,
		},
	}

	return que
}

// 向优先队列中插入一个元素
func (q *PriorityQueue) Offer(t T)  {
	q.a.List = append(q.a.List, t)
	q.swim(q.Size())
}

// 删除并返回最大元素
func (q *PriorityQueue) Poll() (t T) {
	if q.IsEmpty() {
		return
	}
	t = q.a.List[1]
	q.a.List[1] = q.a.List[q.Size()]
	q.a.List = q.a.List[:q.Size()]
	q.sink(1)
	return
}

// 返回最大元素
func (q *PriorityQueue) Peek() (t T) {
	if q.IsEmpty() {
		return
	}
	t = q.a.List[1]
	return
}

// 返回队列是否为空
func (q *PriorityQueue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *PriorityQueue) Size() int {
	return q.a.Len() - 1
}

// 上浮
func (q *PriorityQueue) swim(k int) {
	for k > 1 && q.a.Less(k/2, k) {
		q.a.Swap(k/2, k)
		k /= 2
	}
}

// 下沉
func (q *PriorityQueue) sink(k int) {
	N := q.Size()
	for 2*k <= N {
		j := 2*k
		if j+1 <= N && q.a.Less(j, j+1) {
			j++
		}
		if q.a.Less(j, k) {
			break
		}
		q.a.Swap(j, k)
		k = j
	}
}

func main() {
	//countRestrictedPaths(5, [][]int{{1,2,3},{1,3,3},{2,3,1},{1,4,2},{5,2,2},{3,5,1},{5,4,10}})
	countRestrictedPaths(7, [][]int{{1,3,1},{4,1,2},{7,3,4},{2,5,3},{5,6,1},{6,7,2},{7,5,3},{2,6,4}})
}


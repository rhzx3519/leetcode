/**
@author ZhengHao Lou
Date    2022/07/31
*/
package main

import (
	"fmt"
	"math"
)

const INF = int(math.MaxInt32 >> 1)

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	n := len(edges)
	var adj = map[int][][]int{}
	for i := range edges {
		if edges[i] == -1 {
			continue
		}
		adj[i] = append(adj[i], []int{edges[i], 1})
	}
	d1 := dijkstra(node1, n, adj)
	d2 := dijkstra(node2, n, adj)

	var ans, mx = n, INF
	for i := range edges {
		if d1[i] != INF && d2[i] != INF {
			if max(d1[i], d2[i]) < mx {
				mx = max(d1[i], d2[i])
				ans = i
			} else if max(d1[i], d2[i]) == mx && i < ans {
				ans = i
			}
		}
	}

	if ans == n {
		return -1
	}
	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(ans, mx)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//closestMeetingNode([]int{2, 2, 3, -1}, 0, 1)
	//closestMeetingNode([]int{1, 2, -1}, 0, 2)
	closestMeetingNode([]int{5, 4, 5, 4, 3, 6, -1}, 0, 1)
}

/**
使用优先队列，适用于稀疏图：
想象将所有点分成visited和unvisited两个集合，优先队列记录的是和visited的点相连的边（注意纪录的是边）,队头是这些边中权值最小的一条边。
*/
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
			if dis[ni] > dis[i]+cost {
				dis[ni] = dis[i] + cost
				que.Offer([]int{ni, dis[ni]})
			}
		}
	}

	return dis
}

type (
	T          interface{}
	Comparator func(x, y T) int
)

type Sortable struct {
	List []T
	Cmp  Comparator
}

var (
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
			Cmp:  comparator,
		},
	}

	return que
}

// 向优先队列中插入一个元素
func (q *PriorityQueue) Offer(t T) {
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
		j := 2 * k
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

/**
@author ZhengHao Lou
Date    2021/10/18
*/
package main

import (
	"fmt"
	"math"
)

/**
https://leetcode-cn.com/problems/the-time-when-the-network-becomes-idle/
 */
func networkBecomesIdle(edges [][]int, patience []int) int {
	n := len(patience)
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]int, 0)
	}
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}

	var ans int
	dis := dijkstra(adj, n, 0)
	fmt.Println(dis)
	for i := 1; i < n; i++ {
		round := dis[i] << 1
		p := patience[i]
		if round <= p {
			ans = max(ans, round + 1)
		} else {
			var t int
			if round % p == 0 {
				t = p
			} else {
				t = round % p
			}
			ans = max(ans, round - t + round + 1)
		}
	}
	return ans
}


// 返回x点到达其余节点的最短距离
func dijkstra(adj [][]int, n, x int) []int {
	// initilize
	dis := make([]int, n)
	for i := range dis {
		dis[i] = math.MaxInt32 >> 1
	}
	dis[x] = 0
	que := newPQ()
	que.Offer([]int{x, 0})
	for !que.IsEmpty() {
		p := que.Poll().([]int)
		i, d := p[0], p[1]
		if dis[i] < d {
			continue
		}
		for _, j := range adj[i] {
			if dis[i] + 1 < dis[j] {
				dis[j] = dis[i] + 1
				que.Offer([]int{j, dis[j]})
			}
		}
	}

	//fmt.Println(dis)
	return dis
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func newPQ() *PriorityQueue {
	return &PriorityQueue{
		a: &Sortable{
			List: make([]T, 1),
			Cmp: Reversed(func(x, y T) int {
				a, b := x.([]int)[1], y.([]int)[1]
				if a < b {
					return -1
				} else if a > b {
					return 1
				} else {
					return 0
				}
			}),
		},
	}
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

func New(list []T, comparator Comparator) *PriorityQueue {
	que := &PriorityQueue{
		a: &Sortable{
			List: make([]T, 1),
			Cmp: comparator,
		},
	}

	for _, t := range list {
		que.Offer(t)
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
	fmt.Println(networkBecomesIdle([][]int{{0,1},{1,2}}, []int{0,2,1}))
	fmt.Println(networkBecomesIdle([][]int{{0,1},{0,2},{1,2}}, []int{0,10,10}))
	fmt.Println(networkBecomesIdle([][]int{{5,7},{15,18},{12,6},{5,1},{11,17},{3,9},{6,11},{14,7},{19,13},{13,3},{4,12},{9,15},{2,10},{18,4},{5,14},{17,5},{16,2},{7,1},{0,16},{10,19},{1,8}},
	[]int{0,2,1,1,1,2,2,2,2,1,1,1,2,1,1,1,1,2,1,1}))	// 67
	fmt.Println(networkBecomesIdle([][]int{{3,8},{4,13},{0,7},{0,4},{1,8},{14,1},{7,2},{13,10},{9,11},{12,14},{0,6},{2,12},{11,5},{6,9},{10,3}},
	[]int{0,3,2,1,5,1,5,5,3,1,2,2,2,2,1}))	// 20
}



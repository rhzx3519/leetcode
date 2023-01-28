/*
*
@author ZhengHao Lou
Date    2022/12/13
*/
package main

import "sort"

/*
*
https://leetcode.cn/problems/maximum-number-of-points-from-grid-queries/
*/
var nxt = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func maxPoints(grid [][]int, queries []int) []int {
	m, n := len(grid), len(grid[0])

	var ans []int = make([]int, len(queries))
	var id = make([]int, len(queries))
	for i := range id {
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool {
		return queries[id[i]] < queries[id[j]]
	})
	que := NewPriorityQueue(func(x, y T) int {
		a, b := x.([]int), y.([]int)
		if a[2] < b[2] {
			return 1
		} else if a[2] > b[2] {
			return -1
		}
		return 0
	})
	que.Offer([]int{0, 0, grid[0][0]})
	grid[0][0] = 0

	var c int
	for _, i := range id {
		x := queries[i]
		for !que.IsEmpty() && que.Peek().([]int)[2] < x {
			t := que.Poll().([]int)
			c++
			for _, d := range nxt {
				ni, nj := t[0]+d[0], t[1]+d[1]
				if ni < 0 || ni >= m || nj < 0 || nj >= n || grid[ni][nj] == 0 {
					continue
				}
				que.Offer([]int{ni, nj, grid[ni][nj]})
				grid[ni][nj] = 0
			}
		}
		ans[i] = c
	}

	return ans
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

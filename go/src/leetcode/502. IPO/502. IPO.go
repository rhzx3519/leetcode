package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode-cn.com/problems/ipo/
思路：贪心 + 优先队列
每次选择利润最大的项目，最后的总资本最大
将启动资金 <= w 的所有项目放入优先队列中，优先队列以利润排列，堆顶的利润最大
直到选择了K个项目，或者优先队列中没有可供选择的项目为止
 */
func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	projects := [][]int{}
	for i := range profits {
		projects = append(projects, []int{capital[i], profits[i]})
	}
	n := len(profits)
	sort.Slice(projects, func(i, j int) bool {
		p1, p2 := projects[i], projects[j]
		return p1[0] < p2[0]
	})

	que := &PriorityQueue{
		a: &Sortable{
			List: make([]T, 1),
			Cmp: IntComparator,
		},
	}

	var i int
	for ; k != 0; k-- {
		for i < n && projects[i][0] <= w {
			que.Offer(projects[i][1])
			i++
		}
		if que.IsEmpty() {
			break
		}
		w += que.Poll().(int)
	}

	return w
}

func main() {
	fmt.Println(findMaximizedCapital(2, 0, []int{1,2,3}, []int{0,1,1}))
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

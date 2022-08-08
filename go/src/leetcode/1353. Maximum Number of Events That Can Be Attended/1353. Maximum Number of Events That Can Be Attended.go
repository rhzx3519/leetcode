/**
@author ZhengHao Lou
Date    2021/12/02
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/maximum-number-of-events-that-can-be-attended/
思路：贪心、扫描算法
局部最优：优先选择结束时间早的会议
 */
const N int = 1e5 + 1
func maxEvents(events [][]int) int {
	var left = make([][]int, N)
	for i := range left {
		left[i] = []int{}
	}
	for i, e := range events {
		left[e[0]] = append(left[e[0]], i)
	}

	var count int
	que := NewPriorityQueue(Reversed(IntComparator))
	for i := 1; i < N; i++ {
		// 将开始日期为i的所有event放入pq
		for _, k := range left[i] {
			que.Offer(events[k][1])
		}
		// 将所有结束日期小于i的event删除
		for !que.IsEmpty() && que.Peek().(int) < i {
			que.Poll()
		}

		if !que.IsEmpty() {
			que.Poll()
			count++
		}

	}

	fmt.Println(count)
	return count
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
	maxEvents([][]int{{1,2},{2,3},{3,4}})				// 3
	maxEvents([][]int{{1,2},{2,3},{3,4},{1,2}})				// 4
	maxEvents([][]int{{1,4},{4,4},{2,2},{3,4},{1,1}})   // 4
}
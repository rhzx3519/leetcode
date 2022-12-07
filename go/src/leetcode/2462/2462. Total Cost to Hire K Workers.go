/**
@author ZhengHao Lou
Date    2022/11/07
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode.cn/problems/total-cost-to-hire-k-workers/
*/
func totalCost(costs []int, k int, candidates int) (tot int64) {
	n := len(costs)
	if candidates<<1 >= n {
		sort.Ints(costs)
		for i := 0; i < k; i++ {
			tot += int64(costs[i])
		}
		return
	}

	left, right := NewPriorityQueue(Reversed(IntComparator)), NewPriorityQueue(Reversed(IntComparator))
	for i := 0; i < candidates; i++ {
		left.Offer(costs[i])
		right.Offer(costs[n-1-i])
	}

	var i, j = candidates, n - 1 - candidates
	for ; k != 0; k, n = k-1, n-1 {
		if !left.IsEmpty() && !right.IsEmpty() {
			if left.Peek().(int) <= right.Peek().(int) {
				tot += int64(left.Poll().(int))
				if i <= j {
					left.Offer(costs[i])
					i++
				}

			} else {
				tot += int64(right.Poll().(int))
				if i <= j {
					right.Offer(costs[j])
					j--
				}
			}
		} else if !left.IsEmpty() {
			tot += int64(left.Poll().(int))
		} else {
			tot += int64(right.Poll().(int))
		}

		//fmt.Println(tot)
	}

	return
}

func main() {
	//fmt.Println(totalCost([]int{17, 12, 10, 2, 7, 2, 11, 20, 8}, 3, 4))
	//fmt.Println(totalCost([]int{1, 2, 4, 1}, 3, 3))
	fmt.Println(totalCost([]int{57, 33, 26, 76, 14, 67, 24, 90, 72, 37, 30}, 11, 2))
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

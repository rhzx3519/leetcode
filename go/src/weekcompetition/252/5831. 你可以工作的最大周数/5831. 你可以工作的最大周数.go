package main

import (
	"fmt"
)

func numberOfWeeks(milestones []int) int64 {
	n := len(milestones)
	if n == 1 {
		return 1
	}

	que := New(func(x, y T) int {
		i, j := x.(int), y.(int)
		if milestones[i] < milestones[j] {
			return -1
		} else if milestones[i] > milestones[j] {
			return 1
		} else {
			return 0
		}
	})
	for i := range milestones {
		que.Offer(i)
	}

	var last = -1

	var w int64
	for que.Size() >= 2 {
		i := que.Poll().(int)
		j := que.Poll().(int)
		a, b := milestones[i], milestones[j]
		w += int64(b<<1)
		if last != i && a > b {
			w++
			a -= b + 1
		} else {
			a -= b
		}
		milestones[i], milestones[j] = a, 0
		if a > 0 {
			last = i
			que.Offer(i)
		} else {
			last = -1
		}
	}

	if !que.IsEmpty() {
		i := que.Poll().(int)
		if last != i {
			w++
		}
	}

	fmt.Println( w)
	return w
}


func main() {
	//numberOfWeeks([]int{1,2,3})			// 6
	//numberOfWeeks([]int{5,2,1})			// 7
	//numberOfWeeks([]int{9,3,6,8,2,1})	// 29
	//numberOfWeeks([]int{5,7,5,7,9,7})	// 40
	numberOfWeeks([]int{4,5,5,2})		// 16
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

func New(comparator Comparator) *PriorityQueue {
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

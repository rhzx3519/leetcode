package main

import (
	"fmt"
	"sort"
)

type friend struct {
	id int
	time []int
	seat int
}

func smallestChair(times [][]int, targetFriend int) int {
	friends := []friend{}
	for i, time := range times {
		friends = append(friends, friend{
			id: i,
			time: time,
		})
	}
	sort.Slice(friends, func(i, j int) bool {
		if friends[i].time[0] != friends[j].time[0] {
			return friends[i].time[0] < friends[j].time[0]
		}
		return friends[i].time[1] < friends[j].time[1]
	})

	seats := &PriorityQueue{
		a: &Sortable{
			List: make([]T, 1),
			Cmp: Reversed(IntComparator),
		},
	}
	for i := 0; i < 10001; i++ {
		seats.Offer(i)
	}

	sited := &PriorityQueue{
		a: &Sortable{
			List: make([]T, 1),
			Cmp: func(x, y T) int {
				a, b := x.(friend), y.(friend)
				if a.time[1] < b.time[1] {
					return 1
				} else if a.time[1] > b.time[1] {
					return -1
				} else {
					return 0
				}
			},
		},
	}

	for _, f := range friends {
		for !sited.IsEmpty() && sited.Peek().(friend).time[1] <= f.time[0] {
			tmp := sited.Poll().(friend)
			seats.Offer(tmp.seat)
		}
		f.seat = seats.Poll().(int)
		sited.Offer(f)
		if f.id == targetFriend {
			return f.seat
		}
	}
	return -1
}

func main() {
	fmt.Println(smallestChair([][]int{{1,4},{2,3},{4,6}}, 1))
	fmt.Println(smallestChair([][]int{{3,10},{1,5},{2,6}}, 0))
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

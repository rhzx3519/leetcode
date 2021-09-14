package main

import (
	"fmt"
)

func kthLargestNumber(nums []string, k int) string {
	var numComparator = func(x, y T) int {
		a, b := x.(string), y.(string)
		return compare(a, b)
	}

	que := &PriorityQueue{
		a: &Sortable{
			List: make([]T, 1),
			Cmp: Reversed(numComparator),
		},
	}

	for _, num := range nums {
		que.Offer(num)
		if que.Size() > k {
			que.Poll()
		}
	}
	return que.Peek().(string)
}

func compare(s1, s2 string) int {
	n1, n2 := len(s1), len(s2)
	if n1 != n2 {
		if n1 > n2 {
			return 1
		} else if n1 < n2 {
			return -1
		} else {
			return 0
		}
	}

	for i := 0; i < n1; i++ {
		c1, c2 := s1[i], s2[i]
		if c1 > c2 {
			return 1
		} else if c1 < c2 {
			return -1
		}
	}
	return 0
}

func main() {
	fmt.Println(kthLargestNumber([]string{"3","6","7","10"}, 4))
	fmt.Println(kthLargestNumber([]string{"2","21","12","1"}, 3))
	fmt.Println(kthLargestNumber([]string{"0","0"}, 2))
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
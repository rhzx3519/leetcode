package main

import (
	"fmt"
)

/**
a^c
b^c
e = d^(a^c)^(b^c)^c = d^a^b^c
5^2 = 7
5^1 = 4
6^(5^2)^(5^1)^5

2x2 -> 3x3
11, 2,2
 */
func kthLargestValue(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	pre := make([][]int, m)
	for i := 0; i < m; i++ {
		pre[i] = make([]int, n)
	}
	pre[0][0] = matrix[0][0]

	que := New([]T{pre[0][0]}, IntComparator)

	for i := 1; i < m; i++ {
		pre[i][0] = pre[i-1][0]^matrix[i][0]
		que.Offer(pre[i][0])
		if que.Size() > k {
			que.Poll()
		}
	}
	for j := 1; j < n; j++ {
		pre[0][j] = pre[0][j-1]^matrix[0][j]
		//if _, ok := mapper[pre[0][j]]; ok {
		//	continue
		//}
		que.Offer(pre[0][j])
		if que.Size() > k {
			que.Poll()
		}
	}

	//fmt.Println(pre)

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			pre[i][j] = pre[i-1][j]^pre[i][j-1]^pre[i-1][j-1]^matrix[i][j]
			que.Offer(pre[i][j])
			if que.Size() > k {
				que.Poll()
			}
			//fmt.Println(i, j, pre[i][j])
		}
	}
	//fmt.Println(pre)
	//fmt.Println(que.a)

	return que.Peek().(int)
}

func main() {
	fmt.Println(kthLargestValue([][]int{{5,2},{1,6}}, 1))
	fmt.Println(kthLargestValue([][]int{{5,2},{1,6}}, 2))
	fmt.Println(kthLargestValue([][]int{{5,2},{1,6}}, 3))
	fmt.Println(kthLargestValue([][]int{{5,2},{1,6}}, 4))
}

type (
	T interface {}
	Comparator func(left, right T) int
)

var IntComparator Comparator = func(left, right T) int {
	if left.(int) < right.(int) {
		return 1
	} else if left.(int) > right.(int) {
		return -1
	}
	return 0
}

type Sortable struct {
	List []T
	Cmp Comparator
}
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

/**
@author ZhengHao Lou
Date    2021/11/05
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/maximum-number-of-eaten-apples/
贪心：优先吃先过期的苹果
 */
type app struct {
	i, n, expire int
}

func eatenApples(apples []int, days []int) int {
	n := len(apples)
	que := NewPriorityQueue(func(x, y T) int {
		a, b := x.(app), y.(app)
		if a .expire < b.expire {
			return 1
		} else if a.expire > b.expire {
			return -1
		} else {
			return 0
		}
	})
	var count int
	for i := range apples {
		if apples[i] > 0 {
			que.Offer(app{
				i: i,
				n: apples[i],
				expire: i + days[i],
			})
		}
		for !que.IsEmpty() && que.Peek().(app).expire <= i {
			que.Poll()
		}
		if !que.IsEmpty() {
			a := que.Poll().(app)
			a.n--
			count++
			if a.n > 0 {
				que.Offer(a)
			}
		}
	}

	for day := n; !que.IsEmpty(); day++ {
		for !que.IsEmpty() && que.Peek().(app).expire <= day {
			que.Poll()
		}
		if que.IsEmpty() {
			break
		}
		a := que.Poll().(app)
		//fmt.Println(day, a)
		a.n--
		count++
		if a.n > 0 {
			que.Offer(a)
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
	eatenApples([]int{1,2,3,5,2}, []int{3,2,1,4,2})		// 7
	eatenApples([]int{3,0,0,0,0,2}, []int{3,0,0,0,0,2})	// 5
	eatenApples([]int{2,1,10}, []int{2,10,1})	// 4
	eatenApples([]int{2,1,1,4,5}, []int{10,10,6,4,2})	// 8
}

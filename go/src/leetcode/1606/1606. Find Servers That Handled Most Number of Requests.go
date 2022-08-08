/**
@author ZhengHao Lou
Date    2022/03/31
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/find-servers-that-handled-most-number-of-requests/
思路：模拟
time: O(nlogk)
space: O(n)
*/
func busiestServers(k int, arrival []int, load []int) []int {
	busy := NewPriorityQueue(func(x, y T) int {
		s1, s2 := x.([]int), y.([]int)
		return s2[1] - s1[1]
	})
	idle := make([]int, k)
	for i := 0; i < k; i++ {
		idle[i] = i
	}

	requests := make([]int, k)

	for i, t := range arrival {
		for !busy.IsEmpty() && busy.Peek().([]int)[1] <= t {
			server := busy.Poll().([]int)
			j := lowerBound(idle, server[0])
			idle = append(idle, 0)
			copy(idle[j+1:], idle[j:])
			idle[j] = server[0]
		}
		if len(idle) == 0 {
			continue
		}
		j := lowerBound(idle, i%k)
		if j == len(idle) {
			j = 0
		}
		s := idle[j]
		copy(idle[j:], idle[j+1:])
		idle = idle[:len(idle)-1]
		busy.Offer([]int{s, t + load[i]})
		requests[s]++
	}
	var ans []int
	var mx int
	for i := range requests {
		if requests[i] > mx {
			mx = requests[i]
			ans = []int{i}
		} else if requests[i] == mx {
			ans = append(ans, i)
		}
	}
	fmt.Println(ans)
	return ans
}

func main() {
	busiestServers(3, []int{1, 2, 3, 4, 5}, []int{5, 2, 3, 3, 3})
	busiestServers(3, []int{1, 2, 3, 4}, []int{1, 2, 1, 2})
	busiestServers(3, []int{1, 2, 3}, []int{10, 12, 11})
	busiestServers(3, []int{1, 2, 3, 4, 8, 9, 10}, []int{5, 2, 10, 3, 1, 2, 2})
	busiestServers(1, []int{1}, []int{1})
}

func lowerBound(arr []int, x int) int {
	l, r := 0, len(arr)
	var m int
	for l < r {
		m = l + (r-l)>>1
		if arr[m] >= x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
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

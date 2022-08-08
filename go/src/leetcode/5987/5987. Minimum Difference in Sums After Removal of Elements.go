/**
@author ZhengHao Lou
Date    2022/02/07
*/
package main

import (
	"fmt"
	"math"
)

/**
https://leetcode-cn.com/problems/minimum-difference-in-sums-after-removal-of-elements/
思路：前缀和，后缀和, 动态规划
suf[i]表示num[i:]的最大后缀和
*/
func minimumDifference(nums []int) int64 {
	n := len(nums) / 3

	suf := make([]int64, n*3+1)
	minpq := NewPriorityQueue(Reversed(IntComparator))
	var c int64
	for i := 3*n - 1; i >= n; i-- {
		minpq.Offer(nums[i])
		c += int64(nums[i])
		if minpq.Size() > n {
			t := minpq.Poll().(int)
			c -= int64(t)
		}
		suf[i] = c
	}

	var ans int64 = math.MaxInt64
	maxpq := NewPriorityQueue(IntComparator)
	var pre = int64(0)
	for i := 0; i < 2*n; i++ {
		maxpq.Offer(nums[i])
		pre += int64(nums[i])
		if maxpq.Size() > n {
			t := maxpq.Poll().(int)
			pre -= int64(t)
		}

		if i >= n-1 {
			ans = min(ans, pre-suf[i+1])
		}
	}

	fmt.Println(ans)
	return ans
}

func main() {
	// -1
	minimumDifference([]int{3, 1, 2})
	// 1
	minimumDifference([]int{7, 9, 5, 8, 1, 3})
	// -14
	minimumDifference([]int{16, 46, 43, 41, 42, 14, 36, 49, 50, 28, 38, 25, 17, 5, 18, 11, 14, 21, 23, 39, 23})
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
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

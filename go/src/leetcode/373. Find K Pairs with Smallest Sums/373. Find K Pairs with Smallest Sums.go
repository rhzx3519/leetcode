/**
@author ZhengHao Lou
Date    2022/01/14
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/find-k-pairs-with-smallest-sums/
思路：多路归并，前K对和最小的数对，它的范围肯定在(0,0),(0,1)...(0,n2)..(k,0),(k,1)...(k,n2)之中
 */
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	n1, n2 := len(nums1), len(nums2)
	que := NewPriorityQueue(func(x, y T) int {
		a, b := x.([]int), y.([]int)
		c := (nums1[a[0]] + nums2[a[1]]) - (nums1[b[0]] + nums2[b[1]])
		if c < 0 {
			return 1
		} else if c > 0 {
			return -1
		}
		return 0
	})

	for i := 0; i < min(n1, k); i++ {
		que.Offer([]int{i, 0})
	}

	var ans = [][]int{}
	for ; k != 0 && !que.IsEmpty(); k-- {
		top := que.Poll().([]int)
		ans = append(ans, []int{nums1[top[0]], nums2[top[1]]})
		if top[1] + 1 < n2 {
			que.Offer([]int{top[0], top[1] + 1})
		}
	}

	return ans
}

func main() {
	fmt.Println(kSmallestPairs([]int{1,7,11}, []int{2,4,6}, 3))
	fmt.Println(kSmallestPairs([]int{1,1,2}, []int{1,2,3}, 2))
	fmt.Println(kSmallestPairs([]int{1,2}, []int{3}, 3))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
/**
@author ZhengHao Lou
Date    2022/08/22
*/
package main

import (
	"fmt"
	"sort"
)

/**
从简化问题开始
首先考虑本题的简化问题：给定 nn 个非负数 a_1, a_2, \cdots, a_na
1
​
 ,a
2
​
 ,⋯,a
n
​
 ，求第 kk 个最 小 的子序列和。

这是一个经典问题。我们先把所有数从小到大排序，记 (s, i)(s,i) 表示一个总和为 ss，且最后一个元素是第 ii 个元素的子序列。

我们用一个小根堆维护 (s, i)(s,i)，一开始堆中只有一个元素 (a_1, 1)(a
1
​
 ,1)。当我们取出堆顶元素 (s, i)(s,i) 时，我们可以进行以下操作：

把 a_{i + 1}a
i+1
​
  接到这个子序列的后面形成新的子序列，也就是将 (s + a_{i + 1}, i + 1)(s+a
i+1
​
 ,i+1) 放入堆中。
把子序列中的 a_ia
i
​
  直接换成 a_{i + 1}a
i+1
​
 ，也就是将 (s - a_i + a_{i + 1}, i + 1)(s−a
i
​
 +a
i+1
​
 ,i+1) 放入堆中。
第 (k - 1)(k−1) 次取出的 (s, i)(s,i) 中的 ss 就是答案（k = 1k=1 时答案为空集之和，也就是 00）。

这个做法的正确性基于以下事实：

这种方法能不重不漏地生成所有子序列。
每次放进去的数不小于拿出来的数。
这里不予证明，请读者自行思考。

作者：tsreaper
链接：https://leetcode.cn/problems/find-the-k-sum-of-an-array/solution/by-tsreaper-ps7w/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
type pair struct {
	s int64
	i int
}

func kSum(nums []int, k int) int64 {
	var sum int64
	for i := range nums {
		if nums[i] >= 0 {
			sum += int64(nums[i])
		} else {
			nums[i] = -nums[i]
		}
	}

	pq := NewPriorityQueue(func(x, y T) int {
		p1, p2 := x.(pair), y.(pair)
		if p1.s < p2.s {
			return 1
		} else if p1.s > p2.s {
			return -1
		} else {
			if p1.i < p2.i {
				return 1
			} else if p1.i > p2.i {
				return -1
			}
			return 0
		}
	})

	var kth int64
	n := len(nums)
	sort.Ints(nums)
	pq.Offer(pair{int64(nums[0]), 0})
	for i := 2; i <= k; i++ {
		p := pq.Poll().(pair)
		kth = p.s
		if p.i < n-1 {
			pq.Offer(pair{p.s + int64(nums[p.i+1]), p.i + 1})
			pq.Offer(pair{p.s - int64(nums[p.i]) + int64(nums[p.i+1]), p.i + 1})
		}
	}

	return sum - kth
}

func main() {
	fmt.Println(kSum([]int{2, 4, -2}, 5))
	fmt.Println(kSum([]int{1, -2, 3, 4, -10, 12}, 16))
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

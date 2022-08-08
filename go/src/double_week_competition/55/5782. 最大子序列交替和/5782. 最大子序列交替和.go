package main

import (
	"fmt"
	"math"
)
/**
思路：
dp[i][0] 表示以nums[i-1]结尾的，且子序列大小为偶数的, 最大子序列交替和
dp[i][1] 表示以nums[i-1]结尾的，且子序列大小为奇数的, 最大子序列交替和
dp[i][0] = max(dp[j][1] - int64(nums[i-1]), dp[i][0])
dp[i][1] = max(dp[j][0] + int64(nums[i-1]), dp[i][1])

time: O(n^2)
space: O(n)
 */
func maxAlternatingSum(nums []int) int64 {
	n := len(nums)
	//
	//que := &PriorityQueue{
	//	a: &Sortable{
	//		List: make([]T, 1),
	//		Cmp: Int64Comparator,
	//	},
	//}
	//
	//que.Offer(int64(3))
	//que.Offer(int64(1))
	//que.Offer(int64(2))
	//for !que.IsEmpty() {
	//	fmt.Println(que.Poll())
	//}

	que0 := &PriorityQueue{
		a: &Sortable{
			List: make([]T, 1),
			Cmp: Int64Comparator,
		},
	}
	que0.Offer(int64(0))

	que1 := &PriorityQueue{
		a: &Sortable{
			List: make([]T, 1),
			Cmp: Int64Comparator,
		},
	}
	que1.Offer(int64(0))

	dp := make([][]int64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = []int64{0, 0}
	}
	var ans = int64(math.MinInt64)
	for i := 1; i <= n; i++ {
		//for j := 0; j < i; j++ {
		//	dp[i][0] = max(dp[j][1] - int64(nums[i-1]), dp[i][0])
		//	dp[i][1] = max(dp[j][0] + int64(nums[i-1]), dp[i][1])
		//}

		dp[i][0] = max(que1.Peek().(int64) - int64(nums[i-1]), dp[i][0])
		dp[i][1] = max(que0.Peek().(int64) + int64(nums[i-1]), dp[i][1])

		ans = max(ans, dp[i][0])
		ans = max(ans, dp[i][1])
		que0.Offer(dp[i][0])
		que1.Offer(dp[i][1])
	}
	//fmt.Println(dp)
	return ans
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxAlternatingSum([]int{4,2,5,3}))
	fmt.Println(maxAlternatingSum([]int{5,6,7,8}))
	fmt.Println(maxAlternatingSum([]int{6,2,1,2,4,5}))
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

	Int64Comparator = func(x, y T) int {
		a, b := x.(int64), y.(int64)
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
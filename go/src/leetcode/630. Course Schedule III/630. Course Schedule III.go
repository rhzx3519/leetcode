/**
@author ZhengHao Lou
Date    2021/12/14
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode-cn.com/problems/course-schedule-iii/
思路：按照lastDay递增排序，如果lastDay相同，按照持续时间递增排序
---
具体的，我们先根据「结束时间」对 coursescourses 排升序，从前往后考虑每个课程，处理过程中维护一个总时长 sumsum，
对于某个课程 courses[i]courses[i] 而言，根据如果学习该课程，是否满足「最晚完成时间」要求进行分情况讨论：
学习该课程后，满足「最晚完成时间」要求，即 sum + courses[i][0] <= courses[i][1]sum+courses[i][0]<=courses[i][1]，则进行学习；
学习该课程后，不满足「最晚完成时间」要求，此时从过往学习的课程中找出「持续时间」最长的课程进行「回退」操作（这个持续时长最长的课程有可能是当前课程）。

作者：AC_OIer
链接：https://leetcode-cn.com/problems/course-schedule-iii/solution/gong-shui-san-xie-jing-dian-tan-xin-yun-ghii2/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */
func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		if courses[i][1] != courses[j][1] {
			return courses[i][1] < courses[j][1]
		}
		return courses[i][0] < courses[j][0]
	})
	var sum int
	que := NewPriorityQueue(IntComparator)
	for _, course := range courses {
		d, e := course[0], course[1]
		sum += d
		que.Offer(d)
		if sum > e {
			sum -= que.Poll().(int)
		}
	}
	fmt.Println(que.Size())
	return que.Size()
}

func main() {
	scheduleCourse([][]int{{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200}})
	scheduleCourse([][]int{{1, 2}})
	scheduleCourse([][]int{{3, 2}, {4, 3}})
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

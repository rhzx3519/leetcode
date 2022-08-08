package main

import "fmt"

/**
https://leetcode-cn.com/problems/min-cost-to-connect-all-points/
思路：最小生成树，Prime
1. 将点一个一个连接起来，使用 vi 来标记已经连接的点
2. 将每个新加入的点与其他点的距离计算出来，加入优先队列 pq
3. 从优先队列中翻出最小的距离，并且还未连接的点，进行连接
4. 不断连接新的点，直到所有的点都已连接

 */
func minCostConnectPoints(points [][]int) int {
	n := len(points)

	visited := make([]bool, n)

	que := &PriorityQueue{
		a: &Sortable{
			List: make([]T, 1),
			Cmp: func(x, y T) int {
				e1, e2 := x.([]int), y.([]int)
				if e1[0] < e2[0] {
					return 1
				} else if e1[0] > e2[0] {
					return -1
				} else {
					return 0
				}
			},
		},
	}

	que.Offer([]int{0, 0})
	var ans int
	var count = n
	for !que.IsEmpty() && count > 0 {
		e := que.Poll().([]int)
		d, p := e[0], e[1]
		if visited[p] {
			continue
		}
		ans += d
		visited[p] = true
		count--
		for i := 0; i < n; i++ {
			if visited[i] {
				continue
			}
			que.Offer([]int{manhattan(points[p], points[i]), i})
		}
	}

	return ans
}

func manhattan(p1, p2 []int) int {
	return abs(p1[0] - p2[0]) + abs(p1[1] - p2[1])
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(minCostConnectPoints([][]int{{0,0},{2,2},{3,10},{5,2},{7,0}}))
	fmt.Println(minCostConnectPoints([][]int{{3,12},{-2,5},{-4,1}}))
	fmt.Println(minCostConnectPoints([][]int{{0,0},{1,1},{1,0},{-1,1}}))
	fmt.Println(minCostConnectPoints([][]int{{-1000000,-1000000},{1000000,1000000}}))
	fmt.Println(minCostConnectPoints([][]int{{0,0}}))
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
/**
@author ZhengHao Lou
Date    2022/01/22
*/
package main

import (
	"fmt"
)

type Item struct {
	step, price, row, col int
}

var (
	m, n   int
	INF    int = 1e8
	offset     = [][]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
)

func highestRankedKItems(grid [][]int, pricing []int, start []int, k int) [][]int {
	m, n = len(grid), len(grid[0])
	low, high := pricing[0], pricing[1]

	vis := make([][]int, m)
	for i := range vis {
		vis[i] = make([]int, n)
		for j := range vis[i] {
			vis[i][j] = INF
		}
	}
	vis[start[0]][start[1]] = 0

	pq := NewPriorityQueue(func(x, y T) int {
		it1, it2 := x.(Item), y.(Item)
		if it1.step != it2.step {
			if it1.step > it2.step {
				return 1
			}
			return -1
		} else if it1.price != it2.price {
			if it1.price > it2.price {
				return 1
			}
			return -1
		} else if it1.row != it2.row {
			if it1.row > it2.row {
				return 1
			}
			return -1
		} else if it1.col != it2.col {
			if it1.col > it2.col {
				return 1
			}
			return -1
		}
		return 0
	})

	que := []int{encode(start[0], start[1])}
	for len(que) != 0 {
		t := que[0]
		que = que[1:]
		i, j := decode(t)
		if grid[i][j] >= low && grid[i][j] <= high {
			item := Item{
				step:  vis[i][j],
				price: grid[i][j],
				row:   i,
				col:   j,
			}
			enque(pq, item, k)
		}
		for _, dxy := range offset {
			ni, nj := i+dxy[0], j+dxy[1]
			if ni >= 0 && ni < m && nj >= 0 && nj < n && grid[i][j] != 0 && vis[ni][nj] == INF {
				vis[ni][nj] = vis[i][j] + 1
				que = append(que, encode(ni, nj))
			}
		}
	}

	//fmt.Println(vis)
	var ans = [][]int{}
	for !pq.IsEmpty() {
		item := pq.Poll().(Item)
		//fmt.Println(item.row, item.col)
		ans = append(ans, []int{item.row, item.col})
	}
	reverse(ans)
	return ans
}

func main() {
	// [[0,1],[1,1],[2,1]]
	fmt.Println(highestRankedKItems([][]int{{1, 2, 0, 1}, {1, 3, 0, 1}, {0, 2, 5, 1}}, []int{2, 5}, []int{0, 0}, 3))
	//[[2,1],[1,2]]
	fmt.Println(highestRankedKItems([][]int{{1, 2, 0, 1}, {1, 3, 3, 1}, {0, 2, 5, 1}}, []int{2, 3}, []int{2, 3}, 2))
	// [[1,1],[1,2],[1,0]]
	fmt.Println(highestRankedKItems([][]int{{1, 0, 1}, {3, 5, 2}, {1, 0, 1}}, []int{2, 5}, []int{1, 1}, 9))
}

func reverse(arr [][]int) {
	l, r := 0, len(arr)-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}

func enque(que *PriorityQueue, item Item, k int) {
	que.Offer(item)
	if que.Size() > k {
		que.Poll()
	}
}

func encode(x, y int) int {
	return x*n + y
}

func decode(z int) (int, int) {
	return z / n, z % n
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

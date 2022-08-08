/**
@author ZhengHao Lou
Date    2021/11/03
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/trapping-rain-water-ii/
思路：迭代法
从最外层开始迭代，寻找四周最低的那个格子
time: O(M*N*log(M + N))
space: O(M*N)
 */

type Cell struct {
	h, x, y int
}

var dirs = []int{-1,0,1,0,-1}

func trapRainWater(heightMap [][]int) int {

	que := &PriorityQueue{
		a: &Sortable{
			List: make([]T, 1),
			Cmp: func(x, y T) int {
				c1, c2 := x.(Cell), y.(Cell)
				if c1.h < c2.h {
					return 1
				} else if c1.h > c2.h {
					return -1
				} else {
					return 0
				}
			},
		},
	}

	m, n := len(heightMap), len(heightMap[0])
	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 || i == m-1 || j == n-1 {
				vis[i][j] = true
				que.Offer(Cell{h: heightMap[i][j], x: i, y: j})
			}
		}
	}

	var ans int
	for !que.IsEmpty() {
		cell := que.Poll().(Cell)
		for k := 0; k < 4; k++ {
			nx, ny := cell.x + dirs[k], cell.y + dirs[k+1]
			if nx < 0 || nx >= m-1 || ny < 0 || ny >= n-1 || vis[nx][ny] {
				continue
			}
			vis[nx][ny] = true
			if heightMap[nx][ny] < cell.h {
				ans += cell.h - heightMap[nx][ny]
			}
			que.Offer(Cell{
				h: max(cell.h, heightMap[nx][ny]),
				x: nx,
				y: ny,
			})
		}
	}

	fmt.Println(ans)
	return ans
}



func max(a, b int) int {
	if a > b {
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

func main() {
	trapRainWater([][]int {{1,4,3,1,3,2},{3,2,1,3,2,4},{2,3,3,2,3,1}})	// 4
	trapRainWater([][]int {{3,3,3,3,3},{3,2,2,2,3},{3,2,1,2,3},{3,2,2,2,3},{3,3,3,3,3}})	// 10
	trapRainWater([][]int {{12,13,1,12},{13,4,13,12},{13,8,10,12},{12,13,12,12},{13,13,13,13}})	// 14
}

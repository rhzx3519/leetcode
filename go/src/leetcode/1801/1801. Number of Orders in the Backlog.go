/*
*
@author ZhengHao Lou
Date    2023/01/02
*/
package main

/*
*
https://leetcode.cn/problems/number-of-orders-in-the-backlog/
*/
const mod int = 1e9 + 7

type tuple struct {
	price, amount, orderType int
}

func getNumberOfBacklogOrders(orders [][]int) (tot int) {
	var cmp = func(x, y T) int {
		a, b := x.(*tuple), y.(*tuple)
		if a.orderType == 0 {
			if a.price > b.price {
				return 1
			} else if a.price < b.price {
				return -1
			}
			return 0
		}

		if a.price < b.price {
			return 1
		} else if a.price > b.price {
			return -1
		}
		return 0
	}
	var sell, buy = NewPriorityQueue(cmp), NewPriorityQueue(cmp)

	var match = func() {
		for !sell.IsEmpty() && !buy.IsEmpty() {
			s, b := sell.Peek().(*tuple), buy.Peek().(*tuple)
			if b.price < s.price {
				break
			}

			d := min(s.amount, b.amount)
			s.amount -= d
			b.amount -= d
			if s.amount == 0 {
				sell.Poll()
			} else {
				buy.Poll()
			}
		}
	}

	for i := range orders {
		price, amount, orderType := orders[i][0], orders[i][1], orders[i][2]
		if orderType == 0 {
			buy.Offer(&tuple{price, amount, orderType})
		} else {
			sell.Offer(&tuple{price, amount, orderType})
		}
		match()
	}

	for !sell.IsEmpty() {
		tot = (tot + sell.Poll().(*tuple).amount) % mod
	}
	for !buy.IsEmpty() {
		tot = (tot + buy.Poll().(*tuple).amount) % mod
	}

	return
}

func main() {
	getNumberOfBacklogOrders([][]int{{10, 5, 0}, {15, 2, 1}, {25, 1, 1}, {30, 4, 0}})
	//getNumberOfBacklogOrders([][]int{{7, 1000000000, 1}, {15, 3, 0}, {5, 999999995, 0}, {5, 1, 1}})
	//getNumberOfBacklogOrders([][]int{{1, 29, 1}, {22, 7, 1}, {24, 1, 0}, {25, 15, 1}, {18, 8, 1}, {8, 22, 0}, {25, 15, 1}, {30, 1, 1}, {27, 30, 0}})

}

func min(a, b int) int {
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

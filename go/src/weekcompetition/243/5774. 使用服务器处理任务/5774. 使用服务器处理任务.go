package main

import "fmt"

type server struct {
	i int
	w int
	t_start int
	t_run int
}

func assignTasks(servers []int, tasks []int) []int {
	idles := New([]T{}, func(left, right T) int {
		s1, s2 := left.(server), right.(server)
		if s1.w != s2.w {
			if s1.w < s2.w {
				return 1
			} else if s1.w > s2.w {
				return -1
			}
			return 0
		}
		if s1.i < s2.i {
			return 1
		} else if s1.i > s2.i {
			return -1
		} else {
			return 0
		}
	})

	busies := New([]T{}, func(left, right T) int {
		s1, s2 := left.(server), right.(server)
		if s1.t_start + s1.t_run < s2.t_start + s2.t_run {
			return 1
		} else if s1.t_start + s1.t_run > s2.t_start + s2.t_run {
			return -1
		} else {
			return 0
		}
	})

	ans := make([]int, len(tasks))

	for i, w := range servers {
		idles.Offer(server{i: i, w: w})
	}
	//fmt.Println(idles.a)

	var finished int
    // current time
	var j int

    for cur := 0; finished != len(tasks); {
		for !busies.IsEmpty() {
			serv := busies.Peek().(server)
			if serv.t_start + serv.t_run <= cur {
				busies.Poll()
				idles.Offer(serv)
				finished++
			} else {
				break
			}
		}

		for ; j < len(tasks) && j <= cur && !idles.IsEmpty(); j++ {
			serv := idles.Poll().(server)
			ans[j] = serv.i
			serv.t_start = cur
			serv.t_run = tasks[j]
			busies.Offer(serv)
		}

		if !busies.IsEmpty() && idles.IsEmpty() {
			serv := busies.Peek().(server)
			cur = serv.t_start + serv.t_run
		} else {
			cur++
		}
	}

	fmt.Println(ans)
	return ans
}

func main() {
	assignTasks([]int{3,3,2}, []int{1,2,3,2,1,2})
	assignTasks([]int{5,1,4,3,2}, []int{2,1,2,4,5,2,1})
	assignTasks([]int{10,63,95,16,85,57,83,95,6,29,71}, []int{70,31,83,15,32,67,98,65,56,48,38,90,5})
}

type (
	T interface {}
	Comparator func(left, right T) int
)

type Sortable struct {
	List []T
	Cmp Comparator
}

var (
	IntComparator = func(left, right T) int {
		a, b := left.(int), right.(int)
		if a < b {
			return -1
		} else if a > b {
			return 1
		} else {
			return 0
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
package main

import (
	"fmt"
	"sort"
)

func getSkyline(buildings [][]int) [][]int {
	all := [][]int{}
	for _, b := range buildings {
		all = append(all, []int{b[0], -b[2]})
		all = append(all, []int{b[1], b[2]})
	}

	sort.Slice(all, func(i, j int) bool {
		if all[i][0] != all[j][0] {
			return all[i][0] < all[j][0]
		}
		return all[i][1] < all[j][1]
	})

	ans := [][]int{}
	last := []int{0, 0}
	heights := NewMultiSet()
	heights.Insert(0)
	for _, p := range all {
		if p[1] < 0 {
			heights.Insert(-p[1])
		} else {
			heights.Erase(p[1])
		}

		maxHeight := heights.RBegin()
		if last[1] != maxHeight {
			last[0] = p[0]
			last[1] = maxHeight
			ans = append(ans, append([]int{}, last...))
		}
	}

	return ans
}

func main() {
	fmt.Println(getSkyline([][]int{{2,9,10},{3,7,15},{5,12,12},{15,20,10},{19,24,8}}))
	fmt.Println(getSkyline([][]int{{0,2,3},{2,5,3}}))
}

type MultiSet struct {
	arr []int
}

func NewMultiSet() *MultiSet {
	return &MultiSet{
		arr: make([]int, 0),
	}
}

func (m *MultiSet) Insert(x int) {
	i := m.UpperBound(x)
	m.arr = append(m.arr,  0)
	copy(m.arr[i+1:], m.arr[i:])
	m.arr[i] = x
}

func (m *MultiSet) LowBound(x int) int {
	l, r := 0, len(m.arr)
	var mid int
	for l < r {
		mid = l + (r-l)>>1
		if m.arr[mid] >= x {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func (m *MultiSet) UpperBound(x int) int {
	l, r := 0, len(m.arr)
	var mid int
	for l < r {
		mid = l + (r-l)>>1
		if m.arr[mid] > x {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func (m *MultiSet) Erase(x int) {
	i := m.LowBound(x)
	if i == m.Size() || m.arr[i] != x {
		return
	}
	copy(m.arr[i:], m.arr[i+1:])
	m.arr = m.arr[:len(m.arr)-1]
}

func (m *MultiSet) RBegin() int {
	if m.IsEmpty() {
		panic("MultiSet is empty")
	}
	return m.arr[m.Size()-1]
}

func (m *MultiSet) Size() int {
	return len(m.arr)
}

func (m *MultiSet) IsEmpty() bool {
	return m.Size() == 0
}



















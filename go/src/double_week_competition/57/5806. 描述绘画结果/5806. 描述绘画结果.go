package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode-cn.com/problems/the-skyline-problem/
思路：天际线问题变形
 */
func splitPainting(segments [][]int) [][]int64 {
	all := [][]int{}
	for _, seg := range segments {
		all = append(all, []int{seg[0], -seg[2]})
		all = append(all, []int{seg[1], seg[2]})
	}
	sort.Slice(all, func(i, j int) bool {
		if all[i][0] != all[j][0] {
			return all[i][0] < all[j][0]
		}
		return all[i][1] > all[j][1]
	})

	var ans = make([][]int64, 0)
	last := []int64{int64(all[0][0]), 0}
	colors := NewMultiSet()
	colors.Insert(0)

	var lastSum int64
	for _, p := range all {
		lastSum = colors.Sum()

		if p[1] < 0 {
			colors.Insert(-p[1])
		} else {
			colors.Erase(p[1])
		}

		sum := colors.Sum()
		if last[0] == int64(p[0]) {
			continue
		}
		if sum != lastSum {
			if lastSum > 0 {
				ans = append(ans, []int64{last[0], int64(p[0]), lastSum})
			}
			last[0] = int64(p[0])
			last[1] = sum
		}
	}
	return ans
}

func main() {
	fmt.Println(splitPainting([][]int{{1,4,5},{4,7,7},{1,7,9}}))	// [[1,4,14],[4,7,16]]
	fmt.Println(splitPainting([][]int{{1,7,9},{6,8,15},{8,10,7}}))	// [[1,6,9],[6,7,24],[7,8,15],[8,10,7]]
	fmt.Println(splitPainting([][]int{{1,4,5},{1,4,7},{4,7,1},{4,7,11}}))	// [[1,4,12],[4,7,12]]

	// [[2,3,10],[3,4,34],[4,9,46],[9,10,61],[10,11,36],[11,12,32],[12,13,35],[13,16,21],[18,19,13]]
	fmt.Println(splitPainting([][]int{{4,16,12},{9,10,15},{18,19,13},{3,13,20},{12,16,3},{2,10,10},{3,11,4},{13,16,6}}))
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

func (m *MultiSet) Sum() int64 {
	var s int64
	for i := range m.arr {
		s += int64(m.arr[i])
	}
	return s
}

func (m *MultiSet) Size() int {
	return len(m.arr)
}

func (m *MultiSet) IsEmpty() bool {
	return m.Size() == 0
}

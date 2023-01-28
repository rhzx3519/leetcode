/**
@author ZhengHao Lou
Date    2023/01/09
*/
package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/**
https://leetcode.cn/problems/time-to-cross-a-bridge/
思路：大模拟题，使用堆来模拟
*/
func findCrossingTime(n int, k int, time [][]int) (cur int) {
	sort.SliceStable(time, func(i, j int) bool {
		a, b := time[i], time[j]
		return a[0]+a[2] < b[0]+b[2]
	})
	waitL, waitR := make(hp, k), hp{}
	for i := range waitL {
		waitL[i].i = k - 1 - i
	}
	workL, workR := hp2{}, hp2{}
	for n != 0 {
		for len(workL) != 0 && workL[0].t <= cur { // 左边完成放箱
			heap.Push(&waitL, heap.Pop(&workL))
		}
		for len(workR) != 0 && workR[0].t <= cur { // 右边完成搬箱
			heap.Push(&waitR, heap.Pop(&workR))
		}
		if len(waitR) != 0 { // 右边过桥
			w := heap.Pop(&waitR).(pair)
			cur += time[w.i][2]
			heap.Push(&workL, pair{w.i, cur + time[w.i][3]}) // 记录左边完成放箱的时间
		} else if len(waitL) != 0 { // 左边过桥
			w := heap.Pop(&waitL).(pair)
			cur += time[w.i][0]
			heap.Push(&workR, pair{w.i, cur + time[w.i][1]}) // 记录右边完成搬箱的时间
			n--
		} else if len(workL) == 0 { // cur过小，使用最小放箱/搬箱时间来更新cur
			cur = workR[0].t
		} else if len(workR) == 0 {
			cur = workL[0].t
		} else {
			cur = min(workL[0].t, workR[0].t)
		}
	}
	
	for len(workR) != 0 { // 剩下所有右边未过河的工人依次过河
		w := heap.Pop(&workR).(pair)
		cur = max(cur, w.t) + time[w.i][2]
	}
	
	return cur
}

func main() {
	//fmt.Println(findCrossingTime(1, 3, [][]int{{1, 1, 2, 1}, {1, 1, 3, 1}, {1, 1, 4, 1}}))
	//fmt.Println(findCrossingTime(3, 2, [][]int{{1, 9, 1, 8}, {10, 10, 10, 10}}))
	// 149
	fmt.Println(findCrossingTime(10, 6, [][]int{{2, 10, 5, 8}, {3, 5, 2, 2}, {5, 8, 10, 10},
		{7, 8, 8, 5}, {5, 6, 6, 10}, {6, 10, 6, 2}}))
	// 115
	fmt.Println(findCrossingTime(9, 6, [][]int{{2, 6, 9, 4}, {4, 8, 7, 5}, {4, 6, 7, 6}, {2, 3, 3, 7},
		{9, 3, 6, 8}, {2, 8, 8, 4}}))
}

type pair struct{ i, t int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].i > h[j].i }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

type hp2 []pair

func (h hp2) Len() int            { return len(h) }
func (h hp2) Less(i, j int) bool  { return h[i].t < h[j].t }
func (h hp2) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp2) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp2) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

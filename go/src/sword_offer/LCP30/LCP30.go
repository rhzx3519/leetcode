package main

import (
	"container/heap"
	"fmt"
)

/*
*
https://leetcode.cn/problems/p0NxJO/?envType=daily-question&envId=2024-02-06
*/
func magicTower(nums []int) (tot int) {
	h := hp{}
	heap.Init(&h)
	var hp, delay int64 = 1, 0
	for _, x := range nums {
		if x < 0 {
			heap.Push(&h, tuple{x})
		}
		hp += int64(x)
		if hp <= 0 {
			tot++
			t := heap.Pop(&h).(tuple).val
			hp -= int64(t)
			delay += int64(t)
		}
	}
	hp += delay
	if hp < 0 {
		return -1
	}
	return
}

type tuple struct{ val int }
type hp []tuple

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].val < h[j].val }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func main() {
	fmt.Println(magicTower([]int{100, 100, 100, -250, -60, -140, -50, -50, 100, 150}))
}

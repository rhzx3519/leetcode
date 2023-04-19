package main

import (
	"container/heap"
	"fmt"
	"math"
)

/*
*
https://leetcode.cn/problems/take-gifts-from-the-richest-pile/
*/
func pickGifts(gifts []int, k int) (tot int64) {
	pq := hp{}
	for _, x := range gifts {
		heap.Push(&pq, tuple{val: x})
	}
	for ; k != 0; k-- {
		x := heap.Pop(&pq).(tuple).val
		fmt.Println(x)
		heap.Push(&pq, tuple{val: int(math.Sqrt(float64(x)))})
	}
	for pq.Len() != 0 {
		tot += int64(heap.Pop(&pq).(tuple).val)
	}
	return
}

func main() {
	fmt.Println(pickGifts([]int{25, 64, 9, 4, 100}, 4))
}

type tuple struct{ val int }
type hp []tuple

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].val > h[j].val }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

package main

import (
	"container/heap"
	"fmt"
)

func minStoneSum(piles []int, k int) int {
	var sum int
	pq := hp{}
	for _, p := range piles {
		heap.Push(&pq, tuple{p})
		sum += p
	}
	for ; k != 0; k-- {
		x := heap.Pop(&pq).(tuple).val
		sum -= x >> 1
		heap.Push(&pq, tuple{x - x>>1})
	}
	return sum
}

type tuple struct{ val int }
type hp []tuple

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].val > h[j].val }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func main() {
	fmt.Println(minStoneSum([]int{5, 4, 9}, 2))
	fmt.Println(minStoneSum([]int{4, 3, 6, 7}, 3))
}

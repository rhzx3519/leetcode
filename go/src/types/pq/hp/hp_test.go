/*
*
@author ZhengHao Lou
Date    2022/12/14
*/
package hp

import (
	"container/heap"
	"fmt"
	"testing"
)

type tuple struct{ val, i, j int }
type hp []tuple

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].val < h[j].val }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func TestHP(t *testing.T) {
	h := hp{{1, 0, 0}}
	heap.Init(&h)
	heap.Push(&h, tuple{0, 0, 0})
	heap.Push(&h, tuple{0, 0, 0})
	heap.Push(&h, tuple{2, 0, 0})
	//x := heap.Pop(&h).(tuple)
	//fmt.Println(h)

	for len(h) != 0 {
		fmt.Println(h)
		fmt.Println(heap.Pop(&h).(tuple))
		fmt.Println()
		//if h[0].val != heap.Pop(&h).(tuple).val {
		//}
	}
}

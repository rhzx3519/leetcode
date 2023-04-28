package main

import (
	"container/heap"
	"sort"
)

/*
*
https://leetcode.cn/problems/dinner-plate-stacks/
*/

type DinnerPlates struct {
	capacity int
	stacks   [][]int
	idx      hp
}

func Constructor(capacity int) DinnerPlates {
	return DinnerPlates{capacity: capacity}
}

func (d *DinnerPlates) Push(val int) {
	if d.idx.Len() > 0 && d.idx.IntSlice[0] > len(d.stacks)-1 { // 懒删除
		d.idx = hp{}
	}
	if d.idx.Len() == 0 { // 没有未满栈
		d.stacks = append(d.stacks, []int{val}) // 在末尾添加新栈
		if d.capacity > 1 {
			heap.Push(&d.idx, len(d.stacks)-1)
		}
	} else {
		i := d.idx.IntSlice[0]
		d.stacks[i] = append(d.stacks[i], val)
		if len(d.stacks[i]) == d.capacity { // 栈满了
			heap.Pop(&d.idx)
		}
	}
}

func (d *DinnerPlates) Pop() int {
	return d.PopAtStack(len(d.stacks) - 1)
}

func (d *DinnerPlates) PopAtStack(index int) int {
	if index < 0 || index >= len(d.stacks) || len(d.stacks[index]) == 0 {
		return -1
	}
	if len(d.stacks[index]) == d.capacity { // 满栈
		heap.Push(&d.idx, index) // 出栈就是未满栈了，加入堆中
	}
	bk := len(d.stacks[index]) - 1
	val := d.stacks[index][bk]
	d.stacks[index] = d.stacks[index][:bk]
	for len(d.stacks) > 0 && len(d.stacks[len(d.stacks)-1]) == 0 {
		d.stacks = d.stacks[:len(d.stacks)-1] //去掉末尾的空栈，懒删除，在push时清空堆
	}
	return val
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

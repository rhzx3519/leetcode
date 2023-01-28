/**
@author ZhengHao Lou
Date    2023/01/22
*/
package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/**
https://leetcode.cn/problems/maximum-subsequence-score/
思路：排序 + 最小堆
*/
func maxScore(nums1 []int, nums2 []int, k int) (tot int64) {
	n := len(nums1)
	var ids []int
	for i := 0; i < n; i++ {
		ids = append(ids, i)
	}
	sort.Slice(ids, func(i, j int) bool {
		return nums2[ids[i]] >= nums2[ids[j]]
	})
	pq := hp{}
	var sum int
	for _, i := range ids {
		x, y := nums1[i], nums2[i]
		heap.Push(&pq, tuple{val: x})
		sum += x
		for len(pq) > k {
			sum -= heap.Pop(&pq).(tuple).val
		}
		if len(pq) == k {
			tot = max(tot, int64(y*sum))
		}
	}
	
	return
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

type tuple struct{ val int }
type hp []tuple

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].val < h[j].val }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func main() {
	fmt.Println(maxScore([]int{1, 3, 3, 2}, []int{2, 1, 3, 4}, 3))
	fmt.Println(maxScore([]int{4, 2, 3, 1, 1}, []int{7, 5, 10, 9, 6}, 1))
	fmt.Println(maxScore([]int{2, 1, 14, 12}, []int{11, 7, 13, 6}, 3))
}

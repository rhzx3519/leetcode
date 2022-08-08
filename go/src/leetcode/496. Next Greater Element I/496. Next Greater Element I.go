/**
@author ZhengHao Lou
Date    2021/10/26
*/
package main

import (
	"fmt"
)

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	n2 := len(nums2)
	t := make([]int, n2)
	for i := range nums2 {
		t[i] = -1
	}

	st := []int{}
	for i, num := range nums2 {
		for len(st) > 0 && nums2[st[len(st) - 1]] < num {
			k := st[len(st) - 1]
			st = st[:len(st) - 1]
			t[k] = num
		}
		st = append(st, i)
	}
	fmt.Println(t)

	mapper := make(map[int]int)
	for i, num := range nums2 {
		mapper[num] = i
	}
	next := []int{}
	for _, num := range nums1 {
		i := mapper[num]
		next = append(next, t[i])
	}
	return next
}

func main() {
	fmt.Println(nextGreaterElement([]int{4,1,2}, []int{1,3,4,2}))
	fmt.Println(nextGreaterElement([]int{2,4}, []int{1,2,3,4}))
}

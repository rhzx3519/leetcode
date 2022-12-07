/**
@author ZhengHao Lou
Date    2022/08/13
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode.cn/problems/max-chunks-to-make-sorted-ii/
*/
func maxChunksToSorted(arr []int) (x int) {
	n := len(arr)
	nums := make([]int, n)
	copy(nums, arr)
	sort.Ints(nums)
	var s1, s2 int64
	for i := range nums {
		s1 += int64(nums[i])
		s2 += int64(arr[i])
		if s1 == s2 {
			x++
		}
	}
	return x
}

func main() {
	fmt.Println(maxChunksToSorted([]int{5, 4, 3, 2, 1}))
	//fmt.Println(maxChunksToSorted([]int{2, 1, 3, 4, 4}))
}

package main

import (
	"fmt"
	"sort"
)

/*
*
https://leetcode.cn/problems/find-the-maximum-number-of-marked-indices/
思路：二分 + 贪心
*/
func maxNumOfMarkedIndices(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	return 2 * sort.Search(n>>1, func(k int) bool {
		k++
		for i, x := range nums[:k] {
			if x<<1 > nums[n-k+i] {
				return true
			}
		}
		return false
	})
}

func main() {
	fmt.Println(maxNumOfMarkedIndices([]int{3, 5, 2, 4}))
	fmt.Println(maxNumOfMarkedIndices([]int{9, 2, 5, 4}))
	fmt.Println(maxNumOfMarkedIndices([]int{7, 6, 8}))
}

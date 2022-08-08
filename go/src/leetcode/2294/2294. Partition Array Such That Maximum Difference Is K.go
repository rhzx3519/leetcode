/**
@author ZhengHao Lou
Date    2022/06/06
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode.cn/problems/partition-array-such-that-maximum-difference-is-k/
*/
func partitionArray(nums []int, k int) int {
	var c = 1
	sort.Ints(nums)
	var last = nums[0]
	for i := range nums {
		if nums[i]-last > k {
			c++
			last = nums[i]
		}
	}
	return c
}

func main() {
	fmt.Println(partitionArray([]int{3, 6, 1, 2, 5}, 2))
	fmt.Println(partitionArray([]int{1, 2, 3}, 1))
	fmt.Println(partitionArray([]int{2, 2, 4, 5}, 0))
}

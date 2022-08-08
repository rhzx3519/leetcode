/**
@author ZhengHao Lou
Date    2021/12/03
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode-cn.com/problems/maximize-sum-of-array-after-k-negations/
思路：贪心
 */
func largestSumAfterKNegations(nums []int, k int) int {
	var s int
	n := len(nums)
	var mi = 101
	for i := range nums {
		s += nums[i]
		mi = min(mi, abs(nums[i]))
	}
	sort.Ints(nums)
	//fmt.Println(nums)
	var i int
	for i < n && nums[i] < 0 && k > 0 {
		mi = min(mi, -nums[i])
		s += -nums[i]<<1
		i++
		k--
	}
	if k == 0 {
		return s
	}

	if k&1 == 1 {
		s += -mi<<1
	}

	return s
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(largestSumAfterKNegations([]int{4,2,3}, 1))
	fmt.Println(largestSumAfterKNegations([]int{3,-1,0,2}, 3))
	fmt.Println(largestSumAfterKNegations([]int{2,-3,-1,5,-4}, 2))
	fmt.Println(largestSumAfterKNegations([]int{-2,9,9,8,4}, 5))
	fmt.Println(largestSumAfterKNegations([]int{-8,3,-5,-3,-5,-2}, 6))
	fmt.Println(largestSumAfterKNegations([]int{-4,-2,-3}, 4))
}


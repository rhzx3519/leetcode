/**
@author ZhengHao Lou
Date    2022/05/15
*/
package main

/**
https://leetcode.cn/problems/number-of-ways-to-split-array/
*/
func waysToSplitArray(nums []int) (c int) {
	n := len(nums)
	var s int
	for i := range nums {
		s += nums[i]
	}

	var x int
	for i := 0; i < n-1; i++ {
		x += nums[i]
		if x >= s-x {
			c++
		}
	}
	return c
}

func main() {
	waysToSplitArray([]int{10, 4, -8, 7})
}

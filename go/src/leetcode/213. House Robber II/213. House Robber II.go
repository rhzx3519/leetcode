/**
# 思路：考虑两种情况，由于是环形，拿了第一家就不能拿最后要一家，所以分成
#       两种情况求取最大值就行了，从后往前遍历可以避免f1, f2初始值设置的问题
 */
package main

import "fmt"

func rob(nums []int) int {
	n := len(nums)

	var steal = func(nums []int, l, r int) int {
		d1, d2, d3 := 0, 0, 0
		for i := r; i >= l; i-- {
			d3 = max(d2, d1 + nums[i])
			d1 = d2
			d2 = d3
		}
		return d2
	}

	return max(steal(nums, 0, n-2), steal(nums, 1, n-1))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(rob([]int{2,3,2}))
	fmt.Println(rob([]int{1,2,3,1}))
	fmt.Println(rob([]int{0}))
	fmt.Println(rob([]int{1,3,1,3,100}))
}
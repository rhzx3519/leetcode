/**
@author ZhengHao Lou
Date    2022/07/03
*/
package main

import "math"

/**
https://leetcode.cn/problems/next-greater-element-iii/
*/
func nextGreaterElement(n int) int {
	var nums []int
	for n != 0 {
		nums = append(nums, n%10)
		n /= 10
	}
	reverse(nums)
	n = len(nums)
	var i = n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i < 0 {
		return -1
	}
	var j = n - 1
	for i < j && nums[j] <= nums[i] {
		j--
	}
	if j <= i {
		return -1
	}
	swap(nums, i, j)
	for l, r := i+1, n-1; l < r; l, r = l+1, r-1 {
		swap(nums, l, r)
	}
	var x int64
	for i := range nums {
		x = x*10 + int64(nums[i])
		if x > math.MaxInt32 {
			return -1
		}
	}
	return int(x)
}

func reverse(nums []int) {
	l, r := 0, len(nums)-1
	for l < r {
		swap(nums, l, r)
		l++
		r--
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func main() {
	//nextGreaterElement(12)
	//nextGreaterElement(21)
	nextGreaterElement(2147483486)
}

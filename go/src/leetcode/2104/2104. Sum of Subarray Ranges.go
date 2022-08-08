/**
@author ZhengHao Lou
Date    2022/03/04
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/sum-of-subarray-ranges/
*/
func subArrayRanges(nums []int) int64 {
	n := len(nums)
	mx, mi := section(nums, false), section(nums, true)
	var c int64
	for i := 0; i < n; i++ {
		c += int64(mx[i]-mi[i]) * int64(nums[i])
	}

	fmt.Println(c)
	return c
}

func section(nums []int, isMin bool) []int {
	n := len(nums)
	var st []int
	a := make([]int, n) // a[i]表示以nums[i]为最大值的左边界
	b := make([]int, n) // b[i]表示以nums[i]为最大值的右边界
	for i := range nums {
		for len(st) > 0 && ((isMin && nums[st[len(st)-1]] >= nums[i]) || (!isMin && nums[st[len(st)-1]] <= nums[i])) {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			a[i] = -1
		} else {
			a[i] = st[len(st)-1]
		}
		st = append(st, i)
	}

	st = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && ((isMin && nums[st[len(st)-1]] > nums[i]) || (!isMin && nums[st[len(st)-1]] < nums[i])) {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			b[i] = n
		} else {
			b[i] = st[len(st)-1]
		}
		st = append(st, i)
	}

	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = (i - a[i]) * (b[i] - i)
	}
	fmt.Println(a, b, arr)
	return arr
}

func main() {
	subArrayRanges([]int{1, 3, 3})
}

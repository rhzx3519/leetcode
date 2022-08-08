/**
@author ZhengHao Lou
Date    2022/07/10
*/
package main

/**
https://leetcode.cn/problems/subarray-with-elements-greater-than-varying-threshold/
思路：单调栈
考虑最小单位的子数组nums[i]，以nums[i]为最小值的子数组sub，如果nums[i] > threshold/len(sub)，则为答案；
通过单调栈可以找到左右第一个比nums[i]小的元素位置[l, r],
*/
func validSubarraySize(nums []int, threshold int) int {
	n := len(nums)
	var left, right = make([]int, n), make([]int, n)
	var st []int
	for i := range nums {
		for len(st) != 0 && nums[st[len(st)-1]] >= nums[i] { // left[i] 为左侧小于 nums[i] 的最近元素位置（不存在时为 -1）
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			left[i] = st[len(st)-1]
		} else {
			left[i] = -1
		}
		st = append(st, i)
	}
	st = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(st) != 0 && nums[st[len(st)-1]] >= nums[i] { // right[i] 为右侧小于 nums[i] 的最近元素位置（不存在时为 n）
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			right[i] = st[len(st)-1]
		} else {
			right[i] = n
		}
		st = append(st, i)
	}

	for i := range nums {
		k := right[i] - left[i] - 1
		if nums[i] > threshold/k {
			return k
		}
	}
	return -1
}

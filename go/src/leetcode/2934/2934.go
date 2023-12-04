package main

import "fmt"

/*
*
https://leetcode.cn/problems/minimum-operations-to-maximize-last-elements-in-arrays/
思路：分为两种情况考虑
1. 交换最后一个元素
2. 不交换最后一个元素
*/
func minOperations(nums1 []int, nums2 []int) int {
	n := len(nums1)
	x1 := f(nums1, nums2)
	nums1[n-1], nums2[n-1] = nums2[n-1], nums1[n-1]
	x2 := f(nums1, nums2) + 1

	if x1 >= n && x2 >= n {
		return -1
	}
	return min(x1, x2)
}

func f(nums1 []int, nums2 []int) int {
	n := len(nums1)
	var x int
	for i := len(nums1) - 2; i >= 0; i-- {
		if nums1[i] > nums1[n-1] {
			if nums1[i] <= nums2[n-1] && nums2[i] <= nums1[n-1] {
				x++
			} else {
				return n
			}
		}
	}
	//
	for i := n - 2; i >= 0; i-- {
		if nums2[i] > nums2[n-1] {
			if nums2[i] <= nums1[n-1] && nums1[i] <= nums2[n-1] {
				x++
			} else {
				return n
			}
		}
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
	//fmt.Println(minOperations([]int{1, 2, 7}, []int{4, 5, 3}))
	//fmt.Println(minOperations([]int{2, 3, 4, 5, 9}, []int{8, 8, 4, 4, 4}))
	fmt.Println(minOperations([]int{1, 5, 4}, []int{2, 5, 3}))
}

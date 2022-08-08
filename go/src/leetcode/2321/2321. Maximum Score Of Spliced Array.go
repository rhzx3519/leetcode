/**
@author ZhengHao Lou
Date    2022/06/27
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/maximum-score-of-spliced-array/
*/
func maximumsSplicedArray(nums1 []int, nums2 []int) int {
	var s1, s2 int
	n := len(nums1)
	a := make([]int, n)
	b := make([]int, n)
	for i := range nums1 {
		s1 += nums1[i]
		s2 += nums2[i]
		a[i] = nums2[i] - nums1[i]
		b[i] = nums1[i] - nums2[i]
	}

	// 求nums中和最大的子数组
	var mxa, mxb int
	var sa, sb int
	for i := range a {
		sa += a[i]
		sb += b[i]
		if sa < 0 {
			sa = 0
		}
		if sb < 0 {
			sb = 0
		}
		mxa = max(mxa, sa)
		mxb = max(mxb, sb)
	}

	fmt.Println(s1+mxa, s2+mxb)
	return max(s1+mxa, s2+mxb)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maximumsSplicedArray([]int{60, 60, 60}, []int{10, 90, 10}))
	fmt.Println(maximumsSplicedArray([]int{20, 40, 20, 70, 30}, []int{50, 20, 50, 40, 20}))
	fmt.Println(maximumsSplicedArray([]int{7, 11, 13}, []int{1, 1, 1}))
}

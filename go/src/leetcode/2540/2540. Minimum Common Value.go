/**
@author ZhengHao Lou
Date    2023/01/22
*/
package main

/**
https://leetcode.cn/problems/minimum-common-value/
*/
func getCommon(nums1 []int, nums2 []int) int {
	n1, n2 := len(nums1), len(nums2)
	for i, j := 0, 0; i < n1 && j < n2; {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			return nums2[j]
		}
	}
	return -1
}

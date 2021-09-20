package main

import (
	"fmt"
	"math"
)

/**
https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
 */
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	return (findKth(nums1, nums2, (m + n + 1)/2) + findKth(nums1, nums2, (m + n + 2)/2)) *0.5
}

func findKth(nums1 []int, nums2 []int, k int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		return findKth(nums2, nums1, k)
	}
	if m == 0 {
		return float64(nums2[k-1])
	}
	if k == 1 {
		return float64(min(nums1[0], nums2[0]))
	}
	i := min(m, k/2)
	j := min(n, k/2)
	if nums1[i-1] > nums2[j-1] {
		return findKth(nums1, nums2[j:], k - j)
	} else {
		return findKth(nums1[i:], nums2, k - i)
	}
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func main() {
	nums1 := []int{1,2}
	nums2 := []int{3,4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

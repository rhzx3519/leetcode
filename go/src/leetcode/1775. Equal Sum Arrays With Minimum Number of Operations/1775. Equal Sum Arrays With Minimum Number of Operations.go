/**
TODO
 */
package main

import (
	"fmt"
	"sort"
)

func minOperations(nums1 []int, nums2 []int) int {
	n1, n2 := len(nums1), len(nums2)
	s1, s2 := 0, 0
	for i := 0; i < n1; i++ {
		s1 += nums1[i]
	}
	for i := 0; i < n2; i++ {
		s2 += nums2[i]
	}
	if s1 == s2 {
		return 0
	}
	if s1 > s2 {
		return minOperations(nums2, nums1)
	}

	sort.Ints(nums1)
	sort.Ints(nums2)
	i, j := 0, n2-1
	count := 0
	for i < n1 && j >= 0 && s1 < s2 {
		if 6-nums1[i] > nums2[j]-1 {
			s1 += 6 - nums1[i]
			i++
		} else {
			s2 -= nums2[j] - 1
			j--
		}
		count++
	}

	for i < n1 && s1 < s2 {
		s1 += 6 - nums1[i]
		i++
		count++
	}

	for j >= 0 && s1 < s2 {
		s2 -= nums2[j] - 1
		j--
		count++
	}

	if s1 >= s2 {
		return count
	}

	return -1
}

func main() {
	fmt.Println(minOperations([]int{1,2,3,4,5,6}, []int{1,1,2,2,2,2}))
}
/**
@author ZhengHao Lou
Date    2022/10/04
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/number-of-pairs-satisfying-inequality/
nums1[i] - nums2[i] <= nums1[j] - nums2[j] + diff
*/
func numberOfPairs(nums1 []int, nums2 []int, diff int) (tot int64) {
	n := len(nums1)
	nums := make([]int, n)
	for i := range nums {
		nums[i] = nums1[i] - nums2[i]
	}
	
	var arr []int
	for i := range nums {
		j := upperBound(arr, nums[i]+diff)
		tot += int64(j)
		
		j = upperBound(arr, nums[i])
		arr = append(arr, 0)
		copy(arr[j+1:], arr[j:])
		arr[j] = nums[i]
	}
	return
}

func upperBound(nums []int, x int) int {
	l, r := 0, len(nums)
	for l < r {
		m := l + (r-l)>>1
		if nums[m] > x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func main() {
	fmt.Println(numberOfPairs([]int{3, 2, 5}, []int{2, 2, 1}, 1))
}

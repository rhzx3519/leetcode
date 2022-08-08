/**
@author ZhengHao Lou
Date    2021/11/17
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode-cn.com/problems/advantage-shuffle/
思路：贪心，nums1排序，每次从nums1中寻找>nums2[i]的最小数，如果找不到，则取nums1中的最小值
time: O(nlogn)
space: O(n)
 */
func advantageCount(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	var ans = []int{}
	for _, num := range nums2 {
		i := upperBound(nums1, num)
		if i == len(nums1) {
			ans = append(ans, nums1[0])
			nums1 = nums1[1:]
		} else {
			ans = append(ans, nums1[i])
			copy(nums1[i:], nums1[i+1:])
			nums1 = nums1[:len(nums1) - 1]
		}
	}

	fmt.Println(ans)
	return ans
}

func upperBound(nums []int, x int) int {
	l, r := 0, len(nums)
	var m int
	for l < r {
		m = l + (r - l)>>1
		if nums[m] > x {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}

func main() {
	advantageCount([]int{2,7,11,15}, []int{1,10,4,11})
	advantageCount([]int{12,24,8,32}, []int{13,25,32,11})
}
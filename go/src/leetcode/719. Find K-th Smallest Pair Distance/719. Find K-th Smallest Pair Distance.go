/**
@author ZhengHao Lou
Date    2021/12/22
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode-cn.com/problems/find-k-th-smallest-pair-distance/
思路：二分
 */
func smallestDistancePair(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)

	// 计算距离<=d的距离对数量
	var getLess = func(d int) int {
		var l, c int
		for r := range nums {
			for nums[r] - nums[l] > d {
				l++
			}
			c += r - l
		}
		return c
	}

	var l, r, m int
	l, r = 0, nums[n-1] - nums[0] + 1
	for l < r {
		m = l + (r - l)>>1
		if getLess(m) >= k {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func main() {
	fmt.Println(smallestDistancePair([]int{1,3,1}, 1))
}

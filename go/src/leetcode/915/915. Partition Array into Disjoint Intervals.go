/**
@author ZhengHao Lou
Date    2022/10/24
*/
package main

import (
	"fmt"
	"math"
)

/**
https://leetcode.cn/problems/partition-array-into-disjoint-intervals/
*/
func partitionDisjoint(nums []int) int {
	n := len(nums)
	var left, right = make([]int, n), make([]int, n)
	var mi = math.MaxInt32
	for i := n - 1; i >= 0; i-- {
		if nums[i] < mi {
			mi = nums[i]
		}
		right[i] = mi
	}

	var mx int
	for i := range nums {
		if nums[i] > mx {
			mx = nums[i]
		}
		left[i] = mx
	}
	fmt.Println(left, right)

	for i := 0; i < n-1; i++ {
		if left[i] <= right[i+1] {
			return i
		}
	}
	return -1
}

func main() {
	partitionDisjoint([]int{5, 0, 3, 8, 6})
}

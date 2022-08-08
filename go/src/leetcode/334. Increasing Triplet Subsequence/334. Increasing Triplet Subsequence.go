/**
@author ZhengHao Lou
Date    2022/01/12
*/
package main

import (
	"fmt"
	"math"
)

/**
https://leetcode-cn.com/problems/increasing-triplet-subsequence/
 */
func increasingTriplet(nums []int) bool {
	n := len(nums)
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		left[i] = -1
		right[i] = -1
	}
	var mi = math.MaxInt32
	var miIdx int
	for i := range nums {
		if mi < nums[i] {
			left[i] = miIdx
		}
		if nums[i] < mi {
			mi = nums[i]
			miIdx = i
		}
	}
	var mx = math.MinInt32
	var mxIdx int
	for i := n-1; i >= 0; i-- {
		if mx > nums[i] {
			right[i] = mxIdx
		}
		if nums[i] > mx {
			mx = nums[i]
			mxIdx = i
		}
	}

	fmt.Println(left)
	fmt.Println(right)
	for i := 0; i < n; i++ {
		if left[i] != -1 && right[i] != -1 {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	increasingTriplet([]int{1,2,3,4,5})
	increasingTriplet([]int{5,4,3,2,1})
	increasingTriplet([]int{2,1,5,0,4,6})
	increasingTriplet([]int{20,100,10,12,5,13})
}


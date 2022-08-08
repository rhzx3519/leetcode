/**
@author ZhengHao Lou
Date    2022/04/18
*/
package main

import "math"

/**
https://leetcode-cn.com/problems/find-closest-number-to-zero/
*/
func findClosestNumber(nums []int) int {
	var d = math.MaxInt32
	var x int
	for _, num := range nums {
		if abs(num) < d {
			d = abs(num)
			x = num
		} else if abs(num) == d && num > x {
			x = num
		}
	}
	return x
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

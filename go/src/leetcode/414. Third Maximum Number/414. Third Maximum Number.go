/**
@author ZhengHao Lou
Date    2021/10/06
*/
package main

import "math"

/**
https://leetcode-cn.com/problems/third-maximum-number/
 */
func thirdMax(nums []int) int {
	var f1, f2, f3 = math.MinInt64, math.MinInt64, math.MinInt64
	for _, num := range nums {
		if num > f1 {
			f1, f2, f3 = num, f1, f2
		} else if num < f1 && num > f2 {
			f2, f3 = num, f2
		} else if num < f2 && num > f3 {
			f3 = num
		}
	}
	if f3 != math.MinInt64 {
		return f3
	}
	return f1
}



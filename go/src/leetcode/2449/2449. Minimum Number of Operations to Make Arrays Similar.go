/**
@author ZhengHao Lou
Date    2022/10/30
*/
package main

import (
	"math"
	"sort"
)

/**
https://leetcode.cn/problems/minimum-number-of-operations-to-make-arrays-similar/
*/
func makeSimilar(nums []int, target []int) (ans int64) {
	var f = func(a []int) {
		for i := range a {
			if a[i]&1 == 1 {
				a[i] = -a[i]
			}
		}
		sort.Ints(a)
	}
	f(nums)
	f(target)
	for i := range nums {
		ans += int64(math.Abs(float64(target[i] - nums[i])))
	}
	return ans / 4
}

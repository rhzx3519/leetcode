/**
@author ZhengHao Lou
Date    2022/10/31
*/
package main

/**
https://leetcode.cn/problems/average-value-of-even-numbers-that-are-divisible-by-three/
*/
func averageValue(nums []int) int {
	var s, c int
	for _, i := range nums {
		if i&1 == 0 && i%3 == 0 {
			s += i
			c++
		}
	}
	if c == 0 {
		return 0
	}
	return s / c
}

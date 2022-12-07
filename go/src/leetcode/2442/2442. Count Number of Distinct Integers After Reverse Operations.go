/**
@author ZhengHao Lou
Date    2022/10/18
*/
package main

/**
https://leetcode.cn/problems/count-number-of-distinct-integers-after-reverse-operations/
*/
func countDistinctIntegers(nums []int) int {
	var reverse = func(x int) int {
		var c int
		for x != 0 {
			c = c*10 + x%10
			x /= 10
		}
		return c
	}
	var counter = make(map[int]int)
	for i := range nums {
		counter[nums[i]]++
		counter[reverse(nums[i])]++
	}
	return len(counter)
}

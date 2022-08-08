/**
@author ZhengHao Lou
Date    2022/03/19
*/
package main

/**
https://leetcode-cn.com/contest/biweekly-contest-74/problems/divide-array-into-equal-pairs/
*/
func divideArray(nums []int) bool {
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}
	for _, v := range counter {
		if v%2 != 0 {
			return false
		}
	}
	return true
}

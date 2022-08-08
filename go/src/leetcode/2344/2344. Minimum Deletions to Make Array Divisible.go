/**
@author ZhengHao Lou
Date    2022/07/19
*/
package main

import "sort"

/**
https://leetcode.cn/problems/minimum-deletions-to-make-array-divisible/
*/
func minOperations(nums []int, numsDivide []int) int {
	var counter = make(map[int]int)
	for i := range nums {
		counter[nums[i]]++
	}
	var keys []int
	for k := range counter {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	var divide = make(map[int]int)
	for i := range numsDivide {
		divide[numsDivide[i]]++
	}

	var delete int
NEXT:
	for i := range keys {
		for d := range divide {
			if d%keys[i] != 0 {
				delete += counter[keys[i]]
				continue NEXT
			}
		}
		return delete
	}
	return -1
}

func main() {
	minOperations([]int{2, 3, 2, 4, 3}, []int{9, 6, 9, 3, 15})
	minOperations([]int{4, 3, 6}, []int{8, 2, 6, 10})
}

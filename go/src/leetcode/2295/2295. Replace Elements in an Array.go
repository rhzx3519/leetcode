/**
@author ZhengHao Lou
Date    2022/06/06
*/
package main

/**
https://leetcode.cn/problems/replace-elements-in-an-array/
*/
func arrayChange(nums []int, operations [][]int) []int {
	var counter = make(map[int][]int)
	for i := range nums {
		counter[nums[i]] = append(counter[nums[i]], i)
	}
	for i := range operations {
		counter[operations[i][1]] = counter[operations[i][0]]
		delete(counter, operations[i][0])
	}

	var arr = make([]int, len(nums))
	for num := range counter {
		for _, i := range counter[num] {
			arr[i] = num
		}
	}
	return arr
}

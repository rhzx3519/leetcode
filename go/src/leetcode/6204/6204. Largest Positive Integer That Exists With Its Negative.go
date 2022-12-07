/**
@author ZhengHao Lou
Date    2022/10/17
*/
package main

func findMaxK(nums []int) int {
	var counter = make(map[int]int)
	for i := range nums {
		counter[nums[i]]++
	}
	var mx = -1
	for i := range counter {
		if counter[-i] != 0 {
			mx = max(mx, max(i, -i))
		}
	}
	return mx
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

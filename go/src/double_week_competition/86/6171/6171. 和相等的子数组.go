/**
@author ZhengHao Lou
Date    2022/09/03
*/
package main

func findSubarrays(nums []int) bool {
	var counter = make(map[int]int)
	n := len(nums)
	for i := 0; i < n-1; i++ {
		x := nums[i] + nums[i+1]
		if _, ok := counter[x]; ok {
			return true
		}
		counter[x]++
	}
	return false
}

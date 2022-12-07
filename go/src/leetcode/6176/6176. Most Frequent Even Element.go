/**
@author ZhengHao Lou
Date    2022/09/12
*/
package main

/**
https://leetcode.cn/problems/most-frequent-even-element/
*/
func mostFrequentEven(nums []int) int {
	var ans, mx = -1, 0
	counter := make(map[int]int)
	for i := range nums {
		if nums[i]&1 == 0 {
			counter[nums[i]]++
			if counter[nums[i]] > mx {
				mx = counter[nums[i]]
				ans = nums[i]
			} else if counter[nums[i]] == mx && nums[i] < ans {
				ans = nums[i]
			}
		}
	}
	return ans
}

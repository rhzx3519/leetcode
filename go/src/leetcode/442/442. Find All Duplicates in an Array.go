/**
@author ZhengHao Lou
Date    2022/05/08
*/
package main

/**
https://leetcode-cn.com/problems/find-all-duplicates-in-an-array/
*/
func findDuplicates(nums []int) []int {
	var ans []int
	n := len(nums)
	for i := range nums {
		nums[(nums[i]-1)%n] += n
	}
	for i := range nums {
		if nums[i] > n<<1 {
			ans = append(ans, i+1)
		}
	}

	return ans
}

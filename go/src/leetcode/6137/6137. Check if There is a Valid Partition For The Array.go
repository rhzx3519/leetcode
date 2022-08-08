/**
@author ZhengHao Lou
Date    2022/08/08
*/
package main

/**
https://leetcode.cn/problems/check-if-there-is-a-valid-partition-for-the-array/
*/
func validPartition(nums []int) bool {
	n := len(nums)
	f := make([]bool, n)
	f[0] = false
	f[1] = nums[0] == nums[1]
	for i := 2; i < n; i++ {
		if (nums[i] == nums[i-1]) && f[i-2] { // 1
			f[i] = true
		}

		if (nums[i] == nums[i-1]) && (nums[i-2] == nums[i-1]) && (i == 2 || f[i-3]) {
			f[i] = true
		}
		if (nums[i] == nums[i-1]+1) && (nums[i-2]+1 == nums[i-1]) && (i == 2 || f[i-3]) {
			f[i] = true
		}
	}
	return f[n-1]
}

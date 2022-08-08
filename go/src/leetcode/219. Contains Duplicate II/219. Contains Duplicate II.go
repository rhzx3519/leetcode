/**
@author ZhengHao Lou
Date    2022/01/19
*/
package main

/**
https://leetcode-cn.com/problems/contains-duplicate-ii/
 */
func containsNearbyDuplicate(nums []int, k int) bool {
	mp := make(map[int]int)
	for i := range nums {
		if mp[nums[i]] != 0 {
			return true
		}
		if i >= k {
			mp[nums[i-k]]--
		}
		mp[nums[i]]++
	}
	return false
}

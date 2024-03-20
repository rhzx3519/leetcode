package main

/*
*
https://leetcode.cn/problems/maximum-score-of-a-good-subarray/?envType=daily-question&envId=2024-03-19
*/
func maximumScore(nums []int, k int) (tot int) {
	l, r := k, k
	n := len(nums)
	for {
		for l >= 0 && nums[l] >= nums[k] {
			l--
		}
		for r < n && nums[r] >= nums[k] {
			r++
		}
		tot = max(tot, (r-l-1)*nums[k])
		if l < 0 && r == n {
			break
		}
		if l >= 0 && r < n {
			nums[k] = max(nums[l], nums[r])
		} else if l >= 0 {
			nums[k] = nums[l]
		} else {
			nums[k] = nums[r]
		}
	}
	return
}

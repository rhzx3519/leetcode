package main

func buildArray(nums []int) []int {
	n := len(nums)
	var ans = make([]int, n)
	for i := range nums {
		ans[i] = nums[nums[i]]
	}
	return ans
}

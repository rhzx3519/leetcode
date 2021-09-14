package main

func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}

	var count, ans int
	for i := 2; i < n; i++ {
		if nums[i] - nums[i-1] == nums[i-1] - nums[i-2] {
			count++
		} else {
			count = 0
		}
		ans += count
	}

	return ans
}


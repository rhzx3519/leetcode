package main

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	return max(doRob(nums[1:]), doRob(nums[:n-1]))
}

func doRob(nums []int) int {
	var f1, f2 int
	var f3 int
	for i := range nums {
		f3 = max(f1 + nums[i], f2)
		f1 = f2
		f2 = f3
	}
	return f3
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
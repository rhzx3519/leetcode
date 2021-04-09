package main

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	var mid int
	for l < r {
		mid = l + (r-l)/2
		if nums[mid] >nums[r] {
			l = mid + 1
		} else if nums[mid] < nums[l] {
			r = mid
		} else {
			r--
		}
	}
	return nums[l]
}

// 和 I 的做法类似, 都是二分法, 每次进入无序的那部分找出最小值
// 但是由于有重复值的情况, 需要加入 mid 元素等于 hi 元素的情况
// 此时应该将 hi 减 1 防止重复数字是最小元素
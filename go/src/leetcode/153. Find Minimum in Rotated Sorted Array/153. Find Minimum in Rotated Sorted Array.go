package main

func findMin(nums []int) int {
	n := len(nums)
	l, r := 0, n-1
	var mid int
	for l < r {
		mid = l + (r-l)/2
		if nums[mid] > nums[r] { // minimun value is at mid -> r (4,5,6,1,2,3)
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[l]
}
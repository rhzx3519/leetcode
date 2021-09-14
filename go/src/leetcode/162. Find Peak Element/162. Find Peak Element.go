package main

import "fmt"

func findPeakElement(nums []int) int {
	n := len(nums)
	l, r := 0, n-1
	var m int
	for l < r {
		m = l + (r-l)>>1
		if nums[m+1] < nums[m] { // 6,5
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func main() {
	fmt.Println(findPeakElement([]int{1,2,3,1}))
	fmt.Println(findPeakElement([]int{1,2,1,3,5,6,4}))
}

package main

import "fmt"

/**
time: O(nlogn)
space: O(k)
 */
func maxSlidingWindow(nums []int, k int) []int {
	maximums := []int{}
	win := []int{}
	for i := range nums {
		j := lowerBound(win, nums[i])
		win = append(win, 0)
		copy(win[j+1:], win[j:])
		win[j] = nums[i]

		if i >= k {
			j := lowerBound(win, nums[i-k])
			copy(win[j:], win[j+1:])
			win = win[:len(win) - 1]
		}

		if len(win) == k {
			maximums = append(maximums, win[len(win) - 1])
		}
	}
	return maximums
}

func lowerBound(nums []int, x int) int {
	l, r := 0, len(nums)
	var m int
	for l < r {
		m = l + (r - l)>>1
		if nums[m] >= x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1,3,-1,-3,5,3,6,7}, 3))
}








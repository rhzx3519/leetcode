package main

import "fmt"

func search(nums []int, target int) bool {
	n := len(nums)
	l, r := 0, n-1
	var mid int
	for l <= r {
		for l < r && nums[l] == nums[l+1] {
			l++
		}
		for l < r && nums[r] == nums[r-1] {
			r--
		}

		mid = l + (r - l)/2
		k := nums[mid]
		if k == target {
			return true
		}
		if k < nums[r] { // mid-r 有序
			if k < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else { // l-mid 有序
			if k > target && target >= nums[l] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}

	return false
}

func main() {
	fmt.Println(search([]int{2,5,6,0,0,1,2}, 0))
	fmt.Println(search([]int{2,5,6,0,0,1,2}, 3))
}

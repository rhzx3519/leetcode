package main

import (
	"fmt"
	"sort"
)

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	var count int
	for i := n-1; i >= 0; i-- {
		l, r := 0, i-1
		for l < r {
			if nums[l] + nums[r] > nums[i] {
				count += r - l
				r--
			} else {
				l++
			}
		}
	}
	fmt.Println(count)
	return count
}

func main() {
	triangleNumber([]int{2,2,3,4})
	triangleNumber([]int{4,2,3,4})
}

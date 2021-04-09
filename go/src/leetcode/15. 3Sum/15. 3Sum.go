package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	ans := [][]int{}
	n := len(nums)
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		l, r := i+1, n-1
		for l < r {
			val := nums[i] + nums[l] + nums[r]
			if val < 0 {
				l++
			} else if val > 0 {
				r--
			} else {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				l++
			}
		}
		for ;i < n-1 && nums[i] == nums[i+1]; i++ {}
	}

	return ans
}

func main() {
	fmt.Println(threeSum([]int{-1,0,1,2,-1,-4}))
	fmt.Println(threeSum([]int{}))
	fmt.Println(threeSum([]int{0}))
}

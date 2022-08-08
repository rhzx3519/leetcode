/**
@author ZhengHao Lou
Date    2022/03/21
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/count-hills-and-valleys-in-an-array/
*/
func countHillValley(nums []int) int {
	var c int
	n := len(nums)
	for i := 0; i < n; {
		var j, k = i - 1, i + 1
		for j >= 0 && nums[i] == nums[j] {
			j--
		}
		for k < n && nums[i] == nums[k] {
			k++
		}
		if j >= 0 && k < n {
			if nums[j] > nums[i] && nums[k] > nums[i] {
				c++
			} else if nums[j] < nums[i] && nums[k] < nums[i] {
				c++
			}
		}
		j = i + 1
		for j < n && nums[i] == nums[j] {
			j++
		}
		i = j
	}

	return c
}

func main() {
	fmt.Println(countHillValley([]int{2, 4, 1, 1, 6, 5}))
	fmt.Println(countHillValley([]int{6, 6, 5, 5, 4, 1}))
}

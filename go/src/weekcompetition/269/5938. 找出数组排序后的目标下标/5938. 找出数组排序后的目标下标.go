/**
@author ZhengHao Lou
Date    2021/11/28
*/
package main

import (
	"fmt"
	"sort"
)

func targetIndices(nums []int, target int) []int {
	sort.Ints(nums)
	i := lowerBound(nums, target)
	if i == len(nums) || nums[i] != target {
		return []int{}
	}
	var ans = []int{i}
	var j = i + 1
	for j < len(nums) && nums[j] == target {
		ans = append(ans, j)
		j++
	}
	return ans
}

func lowerBound(nums []int, target int) int {
	l, r := 0, len(nums)
	var m int
	for l < r {
		m = l + (r - l)>>1
		if nums[m] >= target {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func main() {
	fmt.Println(targetIndices([]int{1,2,5,2,3}, 2))
	fmt.Println(targetIndices([]int{1,2,5,2,3}, 3))
	fmt.Println(targetIndices([]int{1,2,5,2,3}, 5))
	fmt.Println(targetIndices([]int{1,2,5,2,3}, 4))
}
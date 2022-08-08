/**
@author ZhengHao Lou
Date    2022/01/24
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode-cn.com/problems/count-elements-with-strictly-smaller-and-greater-elements/
*/
func countElements(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	var i, j = 1, n - 2
	for i < n && nums[i] == nums[i-1] {
		i++
	}
	for j >= 0 && nums[j] == nums[j+1] {
		j--
	}
	return max(0, j-i+1)

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(countElements([]int{11, 7, 2, 15}))
	fmt.Println(countElements([]int{3, 3, 3, 90}))
}

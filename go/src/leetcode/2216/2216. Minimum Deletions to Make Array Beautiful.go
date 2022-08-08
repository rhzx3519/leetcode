/**
@author ZhengHao Lou
Date    2022/03/28
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/minimum-deletions-to-make-array-beautiful/
思路：贪心
*/
func minDeletion(nums []int) int {
	var c int
	n := len(nums)
	for i := 0; i < n; {
		var j = i + 1
		for j < n && nums[j] == nums[i] {
			j++
		}
		c += j - i - 1
		i = j + 1
	}
	c += (n - c) & 1
	fmt.Println(c)
	return c
}

func main() {
	minDeletion([]int{1, 1, 2, 3, 5})
	minDeletion([]int{1, 1, 2, 2, 3, 3})
}

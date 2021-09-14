package main

import "fmt"

/**
https://leetcode-cn.com/problems/array-nesting/
 */
func arrayNesting(nums []int) int {
	n := len(nums)
	var s int
	vis := make([]bool, n)
	for i := range nums {
		if !vis[i] {
			var j = i
			var t int
			for !vis[j] {
				t++
				vis[j] = true
				j = nums[j]
			}

			s = max(s, t)
		}
	}

	fmt.Println(s)
	return s
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arrayNesting([]int{5,4,0,3,1,6,2})
}

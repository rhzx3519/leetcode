package main

import "fmt"

/*
*
https://leetcode.cn/problems/number-of-visible-people-in-a-queue/?envType=daily-question&envId=2024-01-05
思路，单调栈
*/
func canSeePersonsCount(heights []int) []int {
	var ans = make([]int, len(heights))
	var st []int
	for i := len(heights) - 1; i >= 0; i-- {
		for len(st) != 0 && heights[st[len(st)-1]] < heights[i] {
			ans[i]++
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			ans[i]++
		}
		st = append(st, i)
	}
	return ans
}

func main() {
	fmt.Println(canSeePersonsCount([]int{10, 6, 8, 5, 11, 9}))
}

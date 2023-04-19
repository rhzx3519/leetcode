package main

import "fmt"

/*
*
https://leetcode.cn/problems/find-the-substring-with-maximum-cost/
*/
const N = 26

func maximumCostSubstring(s string, chars string, vals []int) int {
	scores := make([]int, N)
	for i := range scores {
		scores[i] = i + 1
	}
	for i, c := range chars {
		scores[int(c-'a')] = vals[i]
	}
	var sum, ans int
	for _, c := range s {
		sum += scores[int(c-'a')]
		if sum < 0 {
			sum = 0
		}
		if sum > ans {
			ans = sum
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumCostSubstring("adaa", "d", []int{-1000}))
	fmt.Println(maximumCostSubstring("abc", "abc", []int{-1, -1, -1}))
}

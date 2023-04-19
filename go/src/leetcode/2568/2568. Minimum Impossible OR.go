package main

/*
*
https://leetcode.cn/problems/minimum-impossible-or/
*/
func minImpossibleOR(nums []int) int {
	var has = make(map[int]bool)
	for _, x := range nums {
		has[x] = true
	}
	for i := 1; ; i <<= 1 {
		if has[i] {
			return i
		}
	}
}

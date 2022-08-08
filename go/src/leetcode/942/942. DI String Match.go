/**
@author ZhengHao Lou
Date    2022/05/09
*/
package main

/**
https://leetcode.cn/problems/di-string-match/
*/
func diStringMatch(s string) []int {
	n := len(s)

	l, r := 0, n
	var nums []int

	for i := range s {
		switch s[i] {
		case 'I':
			nums = append(nums, l)
			l++
		case 'D':
			nums = append(nums, r)
			r--
		}
	}
	nums = append(nums, l)
	return nums
}

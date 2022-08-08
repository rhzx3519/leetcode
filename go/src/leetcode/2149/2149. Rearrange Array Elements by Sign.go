/**
@author ZhengHao Lou
Date    2022/01/24
*/
package main

/**
https://leetcode-cn.com/problems/rearrange-array-elements-by-sign/
*/
func rearrangeArray(nums []int) []int {
	pos, neg := []int{}, []int{}
	for _, num := range nums {
		if num > 0 {
			pos = append(pos, num)
		} else {
			neg = append(neg, num)
		}
	}
	var ans = []int{}
	for i := 0; i < len(nums)>>1; i++ {
		ans = append(ans, pos[i])
		ans = append(ans, neg[i])
	}
	return ans
}

/**
@author ZhengHao Lou
Date    2022/09/19
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/smallest-subarrays-with-maximum-bitwise-or/
*/
func smallestSubarrays(nums []int) []int {
	n := len(nums)

	var ans = make([]int, n)
	for i, x := range nums {
		ans[i] = 1
		for j := i - 1; j >= 0 && nums[j]|x != nums[j]; j-- {
			nums[j] |= x
			ans[j] = i - j + 1
		}
	}
	fmt.Println(ans)
	return ans
}

func main() {
	smallestSubarrays([]int{1, 0, 2, 1, 3})
	smallestSubarrays([]int{1, 2})
}

/**
@author ZhengHao Lou
Date    2022/07/25
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/number-of-zero-filled-subarrays/
*/
func zeroFilledSubarray(nums []int) int64 {
	var ans int64
	var c int
	for i := range nums {
		if nums[i] == 0 {
			c++
			ans += int64(c)
		} else {
			c = 0
		}
	}
	return ans
}

func main() {
	fmt.Println(zeroFilledSubarray([]int{1, 3, 0, 0, 2, 0, 0, 4}))
	fmt.Println(zeroFilledSubarray([]int{0, 0, 0, 2, 0, 0}))
	fmt.Println(zeroFilledSubarray([]int{2, 10, 2019}))
}

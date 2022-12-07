/**
@author ZhengHao Lou
Date    2022/11/16
*/
package main

/**
https://leetcode.cn/problems/number-of-subarrays-with-lcm-equal-to-k/description/
**/
func subarrayLCM(nums []int, k int) (tot int) {
	for i := range nums {
		var lcm = 1
		for _, x := range nums[i:] {
			lcm = lcm / gcd(lcm, x) * x
			if k%lcm != 0 {
				break
			}
			if lcm == k {
				tot++
			}
		}
	}
	return
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

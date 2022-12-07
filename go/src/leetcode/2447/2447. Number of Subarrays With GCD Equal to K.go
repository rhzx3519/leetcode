/**
@author ZhengHao Lou
Date    2022/10/24
*/
package main

/**
https://leetcode.cn/problems/number-of-subarrays-with-gcd-equal-to-k/
*/
func subarrayGCD(nums []int, k int) (count int) {
	for i := range nums {
		var g int
		for j := i; j < len(nums); j++ {
			g = gcd(g, nums[j])
			if g%k != 0 {
				break
			}
			if g == k {
				count++
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

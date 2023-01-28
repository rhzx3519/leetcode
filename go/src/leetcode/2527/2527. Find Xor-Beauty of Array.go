/**
@author ZhengHao Lou
Date    2023/01/08
*/
package main

/**
https://leetcode.cn/problems/find-xor-beauty-of-array/
(x|y)&z ^ (x|z)&y ^ (y|z)&x
x&z | y&z

*/
func xorBeauty(nums []int) (tot int) {
	for i := range nums {
		tot ^= nums[i]
	}
	
	return
}

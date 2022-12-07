/**
@author ZhengHao Lou
Date    2022/10/04
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/bitwise-xor-of-all-pairings/
*/
func xorAllNums(nums1 []int, nums2 []int) (tot int) {
	n1, n2 := len(nums1), len(nums2)
	c1, c2 := n1+n1+n2, n2+n1+n2
	if c1&1 == 1 {
		for i := range nums1 {
			tot ^= nums1[i]
		}
	}
	if c2&1 == 1 {
		for i := range nums2 {
			tot ^= nums2[i]
		}
	}
	
	fmt.Println(tot)
	return tot
}

func main() {
	xorAllNums([]int{2, 1, 3}, []int{10, 2, 5, 0})
}

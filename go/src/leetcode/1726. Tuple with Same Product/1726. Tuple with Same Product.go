package main

import (
	"fmt"
)

/**
https://leetcode-cn.com/problems/tuple-with-same-product/
 */
func tupleSameProduct(nums []int) int {
	var count int
	n := len(nums)
	var s int

	mapper := make(map[int]int)
	for l := 0; l < n; l++ {
		for r := l+1; r < n; r++ {
			// 统计乘积相同的二元组数量
			s = nums[l] * nums[r]
			mapper[s]++
		}
	}
	for _, c := range mapper {
		count += c * (c - 1)
	}

	return count*4
}

func main() {
	fmt.Println(tupleSameProduct([]int{2,3,4,6}))
	fmt.Println(tupleSameProduct([]int{1,2,4,5,10}))
	fmt.Println(tupleSameProduct([]int{2,3,4,6,8,12}))
	fmt.Println(tupleSameProduct([]int{2,3,5,7}))
}
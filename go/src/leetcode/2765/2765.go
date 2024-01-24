package main

import "fmt"

/*
*
https://leetcode.cn/problems/longest-alternating-subarray/?envType=daily-question&envId=2024-01-23
*/
func alternatingSubarray(nums []int) (tot int) {
	for i := range nums {
		k := 1
		j := i + 1
		for ; j < len(nums) && nums[j]-nums[j-1] == k; j++ {
			k *= -1
			tot = max(tot, j-i+1)
		}
	}
	return
}

func main() {
	fmt.Println(alternatingSubarray([]int{2, 3, 4, 3, 4}))
}

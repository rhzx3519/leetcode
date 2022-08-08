/**
@author ZhengHao Lou
Date    2022/04/25
*/
package main

import "math/rand"

/**
https://leetcode-cn.com/problems/random-pick-index/
*/
type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{
		nums: nums,
	}
}

func (this *Solution) Pick(target int) int {
	var ans int
	var j int
	for i, num := range this.nums {
		if num == target {
			j++
			if rand.Intn(j) == 0 {
				ans = i
			}
		}
	}
	return ans
}

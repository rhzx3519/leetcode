/**
@author ZhengHao Lou
Date    2021/11/22
*/
package main

import "math/rand"

/**
https://leetcode-cn.com/problems/shuffle-an-array/
 */
type Solution struct {
	n int
	nums []int
}


func Constructor(nums []int) Solution {
	return Solution{
		n: len(nums),
		nums: nums,
	}
}


func (this *Solution) Reset() []int {
	return this.nums
}


func (this *Solution) Shuffle() []int {
	shuffle := append([]int{}, this.nums...)
	for i := range this.nums {
		j := rand.Intn(this.n - i)
		shuffle[i], shuffle[i+j] = shuffle[i+j], shuffle[i]
	}
	return shuffle
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

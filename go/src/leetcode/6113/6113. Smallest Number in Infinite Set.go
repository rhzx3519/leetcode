/**
@author ZhengHao Lou
Date    2022/07/11
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/smallest-number-in-infinite-set/
*/
type SmallestInfiniteSet struct {
	nums []int
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	var i, j int
	for j < len(this.nums) {
		if i+1 == this.nums[j] {
			i++
			j++
		} else {
			break
		}
	}
	fmt.Println(i + 1)
	this.nums = append(this.nums, 0)
	copy(this.nums[j+1:], this.nums[j:])
	this.nums[j] = i + 1
	return i + 1
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	i := lowerBound(this.nums, num)
	if i != len(this.nums) && this.nums[i] == num {
		copy(this.nums[i:], this.nums[i+1:])
		this.nums = this.nums[:len(this.nums)-1]
		return
	}
}

func lowerBound(nums []int, x int) int {
	l, r := 0, len(nums)
	for l < r {
		m := l + (r-l)>>1
		if nums[m] >= x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func main() {
	obj := Constructor()
	obj.AddBack(2)
	obj.PopSmallest()
	obj.PopSmallest()
	obj.PopSmallest()
	obj.AddBack(1)
	obj.PopSmallest()
	obj.PopSmallest()
	obj.PopSmallest()
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */

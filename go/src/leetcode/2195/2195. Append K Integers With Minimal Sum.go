/**
@author ZhengHao Lou
Date    2022/03/08
*/
package main

import (
	"fmt"
	"math"
	"sort"
)

/**
https://leetcode-cn.com/problems/append-k-integers-with-minimal-sum/
*/
func minimalKSum(nums []int, k int) int64 {
	var s int64
	sort.Ints(nums)
	nums = append(nums, []int{math.MaxInt32, math.MaxInt32}...)
	copy(nums[1:], nums[0:])
	nums[0] = 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		t := nums[i] - nums[i-1] - 1
		if t < k {
			s += int64(sn(nums[i-1]+1, nums[i-1]+t))
			k -= t
		} else {
			s += int64(sn(nums[i-1]+1, nums[i-1]+k))
			break
		}
	}

	fmt.Println(s)
	return s
}

func sn(a1, an int) int {
	return (an - a1 + 1) * (a1 + an) / 2
}

func main() {
	//minimalKSum([]int{1, 4, 25, 10, 25}, 2)
	//minimalKSum([]int{5, 6}, 6)
	// 794
	minimalKSum([]int{96, 44, 99, 25, 61, 84, 88, 18, 19, 33, 60, 86, 52,
		19, 32, 47, 35, 50, 94, 17, 29, 98, 22, 21, 72, 100, 40, 84}, 35)
}

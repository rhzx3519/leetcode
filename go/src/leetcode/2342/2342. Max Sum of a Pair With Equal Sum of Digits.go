/**
@author ZhengHao Lou
Date    2022/07/19
*/
package main

import "sort"

/**
https://leetcode.cn/problems/max-sum-of-a-pair-with-equal-sum-of-digits/
*/
func maximumSum(nums []int) int {
	var sum = func(x int) int {
		var s int
		for x != 0 {
			s += x % 10
			x /= 10
		}
		return s
	}
	var ans = -1
	var counter = make(map[int][]int)
	for i := range nums {
		s := sum(nums[i])
		counter[s] = append(counter[s], i)
		sort.Slice(counter[s], func(i, j int) bool {
			return nums[counter[s][i]] > nums[counter[s][j]]
		})

		if len(counter[s]) > 2 {
			counter[s] = counter[s][:2]
		}
		if len(counter[s]) == 2 {
			ans = max(ans, nums[counter[s][0]]+nums[counter[s][1]])
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	maximumSum([]int{18, 43, 36, 13, 7})
}

/**
@author ZhengHao Lou
Date    2022/05/01
*/
package main

import "math"

/**
https://leetcode-cn.com/problems/minimum-average-difference/
*/
func minimumAverageDifference(nums []int) int {
	n := len(nums)
	pre := make([]int, n+1)
	for i := range nums {
		pre[i+1] = pre[i] + nums[i]
	}
	
	var mi, ans int = math.MaxInt32, math.MaxInt32
	for i := range nums {
		var front, rear int
		if i < n-1 {
			rear = (pre[n] - pre[i+1]) / (n - i - 1)
		}
		front = pre[i+1] / (i + 1)
		d := abs(front - rear)
		if d < mi {
			mi = d
			ans = i
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	minimumAverageDifference([]int{2, 5, 3, 9, 5, 3})
	minimumAverageDifference([]int{0})
}

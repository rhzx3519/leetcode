/**
@author ZhengHao Lou
Date    2022/11/07
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/maximum-sum-of-distinct-subarrays-with-length-k/
*/
func maximumSubarraySum(nums []int, k int) int64 {
	n := len(nums)
	sums := make([]int64, n+1)
	for i := range nums {
		sums[i+1] = sums[i] + int64(nums[i])
	}

	var counter = make(map[int]int)
	var ans int64
	for i := 0; i < n; i++ {
		if i >= k {
			if sums[i]-sums[i-k] >= ans && len(counter) == k {
				ans = sums[i] - sums[i-k]
			}
			counter[nums[i-k]]--
			if counter[nums[i-k]] == 0 {
				delete(counter, nums[i-k])
			}
		}
		counter[nums[i]]++
	}
	if sums[n]-sums[n-k] >= ans && len(counter) == k {
		ans = sums[n] - sums[n-k]
	}
	return ans
}

func main() {
	//fmt.Println(maximumSubarraySum([]int{1, 5, 4, 2, 9, 9, 9}, 3))
	//fmt.Println(maximumSubarraySum([]int{4, 4, 4}, 3))
	fmt.Println(maximumSubarraySum([]int{1, 1, 1, 7, 8, 9}, 3)) // 24
}

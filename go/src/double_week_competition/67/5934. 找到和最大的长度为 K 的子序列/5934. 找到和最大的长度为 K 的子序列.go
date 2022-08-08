/**
@author ZhengHao Lou
Date    2021/12/11
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode-cn.com/contest/biweekly-contest-67/problems/find-subsequence-of-length-k-with-the-largest-sum/
 */
func maxSubsequence(nums []int, k int) []int {
	var index = []int{}
	for i := range nums {
		index = append(index, i)
	}
	sort.Slice(index, func(i, j int) bool {
		return nums[index[i]] > nums[index[j]]
	})
	hash := map[int]int{}
	for i := 0; i < k; i++ {
		hash[index[i]]++
	}
	var ans = []int{}
	for i := range nums {
		if hash[i] != 0 {
			ans = append(ans, nums[i])
		}
	}

	fmt.Println(ans)
	return ans
}

func main() {
	maxSubsequence([]int{2,1,3,3}, 2)
	maxSubsequence([]int{-1,-2,3,4}, 3)
	maxSubsequence([]int{3,4,3,3}, 2)
	maxSubsequence([]int{50, -75}, 2)
}



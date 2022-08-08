/**
@author ZhengHao Lou
Date    2022/03/14
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/find-all-k-distant-indices-in-an-array/
time: O(n)
space: O(n)
*/
func findKDistantIndices(nums []int, key int, k int) []int {
	var ans []int
	n := len(nums)
	for i, num := range nums {
		if num == key {
			l := max(0, i-k)
			if len(ans) != 0 {
				l = max(l, ans[len(ans)-1]+1)
			}
			for j := l; j <= min(n-1, i+k); j++ {
				ans = append(ans, j)
			}
		}
	}

	fmt.Println(ans)
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	findKDistantIndices([]int{3, 4, 9, 1, 3, 9, 5}, 9, 1)
	findKDistantIndices([]int{2, 2, 2, 2, 2}, 2, 2)
}

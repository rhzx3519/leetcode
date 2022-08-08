/**
@author ZhengHao Lou
Date    2021/11/05
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/longest-arithmetic-subsequence-of-given-difference/
 */
func longestSubsequence(arr []int, difference int) int {
	var ans int

	mapper := make(map[int]int)
	for _, num := range arr {
		if j, ok := mapper[num - difference]; ok {
			mapper[num] = j + 1
		} else {
			mapper[num] = 1
		}
		ans = max(ans, mapper[num])
	}
	fmt.Println(ans)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	longestSubsequence([]int{1,2,3,4}, 1)
	longestSubsequence([]int{1,3,5,7}, 1)
	longestSubsequence([]int{1,5,7,8,5,3,4,2,1}, -2)
	longestSubsequence([]int{3,0,-3,4,-4,7,6}, 3) // 2
}

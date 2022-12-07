/**
@author ZhengHao Lou
Date    2022/08/29
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode.cn/problems/longest-subsequence-with-limited-sum/
*/
func answerQueries(nums []int, queries []int) []int {
	n := len(nums)
	sort.Ints(nums)
	sums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	var ans []int
	for j := range queries {
		k := sort.Search(n+1, func(i int) bool {
			return sums[i] > queries[j]
		})
		ans = append(ans, k-1)
	}

	fmt.Println(ans)
	return ans
}

func main() {
	answerQueries([]int{4, 5, 2, 1}, []int{3, 10, 21})
	answerQueries([]int{4, 5, 2, 3}, []int{1})
}

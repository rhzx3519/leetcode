/**
@author ZhengHao Lou
Date    2022/07/19
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode.cn/problems/query-kth-smallest-trimmed-number/
*/
func smallestTrimmedNumbers(nums []string, queries [][]int) []int {
	var ans = make([]int, len(queries))
	index := make([]int, len(nums))
	for i := range nums {
		index[i] = i
	}

	for i := range queries {
		k, trim := queries[i][0], queries[i][1]
		sort.Slice(index, func(i, j int) bool {
			a, b := nums[index[i]][len(nums[index[i]])-trim:], nums[index[j]][len(nums[index[j]])-trim:]
			if a != b {
				return a < b
			}
			return index[i] < index[j]
		})
		ans[i] = index[k-1]
	}

	fmt.Println(ans)
	return ans
}

func main() {
	smallestTrimmedNumbers([]string{"102", "473", "251", "814"}, [][]int{{1, 1}, {2, 3}, {4, 2}, {1, 2}})
}

package main

import (
	"fmt"
	"slices"
)

/*
*
https://leetcode.cn/problems/make-lexicographically-smallest-array-by-swapping-elements/
思路：分组排序，值相差limit之内的元素可以按照顺序任意排序
*/
func lexicographicallySmallestArray(nums []int, limit int) []int {
	n := len(nums)
	ids := make([]int, n)
	for i := range ids {
		ids[i] = i
	}
	// i, j 表示ids[i], ids[j]
	slices.SortFunc(ids, func(i, j int) int {
		return nums[i] - nums[j]
	})
	ans := make([]int, n)
	for i := 0; i < n; {
		j := i + 1
		for ; j < n && nums[ids[j]]-nums[ids[j-1]] <= limit; j++ {
		}
		subIds := slices.Clone(ids[i:j])
		slices.Sort(subIds)
		for k := range subIds {
			ans[subIds[k]] = nums[ids[i+k]]
		}
		i = j
	}

	return ans
}

func main() {
	fmt.Println(lexicographicallySmallestArray([]int{1, 7, 6, 18, 2, 1}, 3))
}

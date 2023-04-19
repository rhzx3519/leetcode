package main

import (
	"fmt"
	"sort"
)

/*
*
https://leetcode.cn/problems/find-score-of-an-array-after-marking-all-elements/
*/
func findScore(nums []int) (tot int64) {
	n := len(nums)
	ids := make([]int, n)
	for i := range nums {
		ids[i] = i
	}
	sort.Slice(ids, func(i, j int) bool {
		if nums[ids[i]] != nums[ids[j]] {
			return nums[ids[i]] < nums[ids[j]]
		}
		return ids[i] < ids[j]
	})
	vis := make([]bool, n)
	for len(ids) != 0 {
		for len(ids) != 0 && vis[ids[0]] {
			ids = ids[1:]
		}
		if len(ids) == 0 {
			break
		}
		index := ids[0]
		ids = ids[1:]
		x := nums[index]
		tot += int64(x)

		fmt.Println(index, x)
		if index-1 >= 0 {
			vis[index-1] = true
		}
		if index+1 < n {
			vis[index+1] = true
		}
	}

	return
}

func main() {
	//fmt.Println(findScore([]int{2, 1, 3, 4, 5, 2}))
	//fmt.Println(findScore([]int{2, 3, 5, 1, 3, 2}))
	fmt.Println(findScore([]int{5, 5, 6, 8, 10, 5, 1})) //22
}

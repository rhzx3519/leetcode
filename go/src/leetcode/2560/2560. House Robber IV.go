package main

import (
	"fmt"
	"sort"
)

/*
*
https://leetcode.cn/problems/house-robber-iv/
*/
func minCapability(nums []int, k int) int {
	return sort.Search(1e9, func(mx int) bool {
		var f0, f1 int
		for _, x := range nums {
			if x <= mx {
				f0, f1 = f1, max(f1, f0+1)
			} else {
				f0 = f1
			}
		}
		return f1 >= k
	})
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minCapability([]int{2, 3, 5, 9}, 2))
}

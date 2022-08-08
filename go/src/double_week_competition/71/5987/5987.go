/**
@author ZhengHao Lou
Date    2022/02/05
*/
package main

import (
	"fmt"
	"math"
)

func minimumDifference(nums []int) int64 {
	var n = len(nums) / 3

	var diff int64 = math.MaxInt64 >> 1
	var dfs func(idx, c1, c2 int, left, right int64)
	dfs = func(idx, c1, c2 int, left, right int64) {
		if idx == len(nums) {
			if c2 == n {
				if left-right < diff {
					diff = left - right
				}
			}
			return
		}
		if c1 > n || c2 > n {
			return
		}

		if c1 < n {
			dfs(idx+1, c1+1, c2, left+int64(nums[idx]), right)
		} else {
			dfs(idx+1, c1, c2, left, right+int64(nums[idx]))
		}
		dfs(idx+1, c1, c2+1, left, right)
	}

	dfs(0, 0, 0, 0, 0)
	fmt.Println(diff)
	return diff
}

func main() {
	minimumDifference([]int{3, 1, 2})
	minimumDifference([]int{7, 9, 5, 8, 1, 3})
}

/**
@author ZhengHao Lou
Date    2022/08/08
*/
package main

import (
	"fmt"
)

/**
https://leetcode.cn/problems/minimum-replacements-to-sort-the-array/
*/
func minimumReplacement(nums []int) (step int64) {
	n := len(nums)
	var mi = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		x := nums[i]
		if x > mi {
			k := (x + mi - 1) / mi
			step += int64(k - 1)
			mi = min(mi, x/k)
		} else {
			mi = min(mi, x)
		}
	}
	fmt.Println(step)
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	minimumReplacement([]int{3, 9, 3})
	minimumReplacement([]int{1, 2, 3, 4, 5})
	minimumReplacement([]int{12, 9, 7, 6, 17, 19, 21}) // 3,3,3,3, 3,3,3, 3,4, 6, 17, 19, 21
	// 10
	minimumReplacement([]int{7, 6, 15, 6, 11, 14, 10}) // 1,3,3, 3,3, 3,3,3,3,3, 3,3, 5,6, 7,7, 10
}

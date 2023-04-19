package main

import (
	"fmt"
	"sort"
)

/*
*
https://leetcode.cn/problems/arithmetic-subarrays/
思路：预处理
*/
func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	check := func(arr []int) bool {
		if len(arr) == 1 {
			return false
		} else if len(arr) == 2 {
			return true
		}
		sort.Ints(arr)
		fmt.Println(arr)
		for i := 1; i < len(arr)-1; i++ {
			if arr[i]-arr[i-1] != arr[i+1]-arr[i] {
				return false
			}
		}
		return true
	}
	var ans = make([]bool, len(l))
	for i := range ans {
		ans[i] = check(append([]int{}, nums[l[i]:r[i]+1]...))
	}
	return ans
}

func main() {
	checkArithmeticSubarrays([]int{4, 6, 5, 9, 3, 7}, []int{0, 0, 2}, []int{2, 3, 5})
}

/**
@author ZhengHao Lou
Date    2021/11/29
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode-cn.com/problems/k-th-smallest-prime-fraction/
 */
func kthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)
	nums := [][]int{}
	for i := 0; i < n; i++ {
		for j := i+1; j < n; j++ {
			nums = append(nums, []int{arr[i], arr[j]})
		}
	}

	sort.Slice(nums, func(i, j int) bool {
		return float64(nums[i][0]) / float64(nums[i][1]) < float64(nums[j][0]) / float64(nums[j][1])
	})
	return nums[k - 1]
}

func main() {
	fmt.Println(kthSmallestPrimeFraction([]int{1,2,3,5}, 3))
}

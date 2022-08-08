/**
@author ZhengHao Lou
Date    2021/10/10
*/
package main

import (
	"fmt"
	"sort"
)

func minOperations(grid [][]int, x int) int {
	nums := []int{}
	for i := range grid {
		for j := range grid[i] {
			nums = append(nums, grid[i][j])
		}
	}
	sort.Ints(nums)
	n := len(nums)
	var m = nums[n >> 1]
	//if n & 1 == 1 { // odd
	//	mid = nums[n>>1]
	//} else {
	//	mid = nums
	//}

	var count int
	for _, num := range nums {
		if abs(m - num) % x != 0 {
			return -1
		}
		count += abs(m - num) / x
	}
	return count
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func main() {
	fmt.Println(minOperations([][]int{{2,4},{6,8}}, 2))
	fmt.Println(minOperations([][]int{{1,5},{2,3}}, 1))
	fmt.Println(minOperations([][]int{{1,2},{3,4}}, 2))
}

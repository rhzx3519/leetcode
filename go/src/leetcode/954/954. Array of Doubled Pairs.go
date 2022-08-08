/**
@author ZhengHao Lou
Date    2022/04/01
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode-cn.com/problems/array-of-doubled-pairs/
*/
func canReorderDoubled(arr []int) bool {
	n := len(arr)
	counter := make(map[int]int)
	for _, num := range arr {
		counter[num]++
	}

	var nums []int
	for num := range counter {
		nums = append(nums, num)
	}
	sort.Ints(nums)

	var c int
	for _, num := range nums {
		if num == 0 {
			c += counter[0] >> 1
			counter[0] &= 1
		} else {
			mi := min(counter[num], counter[num<<1])
			c += mi
			counter[num] -= mi
			counter[num<<1] -= mi
		}
	}

	fmt.Println(c)
	return c >= n>>1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	canReorderDoubled([]int{4, -2, 2, -4})
}

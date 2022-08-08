/**
@author ZhengHao Lou
Date    2022/07/26
*/
package main

import (
	"fmt"
	"math/bits"
)

/**
https://leetcode.cn/problems/number-of-excellent-pairs/
*/
const N = 30

func countExcellentPairs(nums []int, k int) int64 {
	counter := make([]int, N+1)
	st := make(map[int]int)
	for i := range nums {
		st[nums[i]]++
	}
	for x := range st {
		counter[bits.OnesCount32(uint32(x))]++
	}
	
	var ans int
	for i := range counter {
		for j := range counter {
			if i+j >= k {
				ans += counter[i] * counter[j]
			}
		}
	}
	
	return int64(ans)
}

func main() {
	fmt.Println(countExcellentPairs([]int{1, 2, 3, 1}, 3))
}

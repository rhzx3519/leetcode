/**
@author ZhengHao Lou
Date    2021/11/22
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode-cn.com/problems/array-of-doubled-pairs/
思路：求n/2对 二倍数对
time: O(nlogn)
space: O(n)
 */
func canReorderDoubled(arr []int) bool {
	n := len(arr)
	counter := make(map[int]int)
	for _, num := range arr {
		counter[num]++
	}
	keys := []int{}
	for k, _ := range counter {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	// n/2对 二倍数对
	var count int
	for _, k := range keys {
		if k == 0 {
			if counter[k] >= 2 {
				count += counter[k] >> 1
				counter[k] &= 1
			}
		} else if counter[k<<1] != 0 {
			minC := min(counter[k], counter[k<<1])
			count += minC
			counter[k] -= minC
			counter[k<<1] -= minC
		}
	}
	return count >= n>>1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	//fmt.Println(canReorderDoubled([]int{3,1,3,6}))
	//fmt.Println(canReorderDoubled([]int{2,1,2,6}))
	//fmt.Println(canReorderDoubled([]int{4,-2,2,-4}))
	//fmt.Println(canReorderDoubled([]int{1,2,4,16,8,4}))
	fmt.Println(canReorderDoubled([]int{-33,0}))		// false
	fmt.Println(canReorderDoubled([]int{2,4,0,0,8,1}))	// true
}

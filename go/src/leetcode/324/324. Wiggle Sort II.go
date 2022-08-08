/**
@author ZhengHao Lou
Date    2022/06/28
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode.cn/problems/wiggle-sort-ii/
*/
func wiggleSort(nums []int) {
	n := len(nums)
	sort.Ints(nums)
	tmp := append([]int{}, nums...)
	j, k := (n+1)>>1, n
	for i := range tmp {
		if i&1 == 0 {
			j--
			nums[i] = tmp[j]
		} else {
			k--
			nums[i] = tmp[k]
		}
	}
	fmt.Println(nums)
}

func main() {
	wiggleSort([]int{1, 5, 1, 1, 6, 4})
	wiggleSort([]int{1, 3, 2, 2, 3, 1}) // 1,1,2,2,3,3, -> 2,3,2,3,1,1
}

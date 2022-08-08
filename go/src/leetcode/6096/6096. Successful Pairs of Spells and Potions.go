/**
@author ZhengHao Lou
Date    2022/06/12
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode.cn/problems/successful-pairs-of-spells-and-potions/
*/
func successfulPairs(spells []int, potions []int, success int64) []int {
	_, m := len(spells), len(potions)
	sort.Ints(potions)

	var ans []int
	for i := range spells {
		e := success / int64(spells[i])
		if success%int64(spells[i]) != 0 {
			e++
		}
		j := lowerBound(potions, int(e))
		ans = append(ans, m-j)
	}
	return ans
}

func lowerBound(nums []int, x int) int {
	l, r := 0, len(nums)
	var m int
	for l < r {
		m = l + (r-l)>>1
		if nums[m] >= x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func main() {
	fmt.Println(successfulPairs([]int{5, 1, 3}, []int{1, 2, 3, 4, 5}, 7))
}

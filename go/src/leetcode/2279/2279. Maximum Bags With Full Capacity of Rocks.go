/**
@author ZhengHao Lou
Date    2022/05/23
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode.cn/problems/maximum-bags-with-full-capacity-of-rocks/
*/
func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	var bags []int
	for i := range capacity {
		bags = append(bags, i)
	}
	sort.Slice(bags, func(i, j int) bool {
		return capacity[bags[i]]-rocks[bags[i]] < capacity[bags[j]]-rocks[bags[j]]
	})
	for i := range bags {
		c := min(capacity[bags[i]]-rocks[bags[i]], additionalRocks)
		additionalRocks -= c
		rocks[bags[i]] += c
		if additionalRocks <= 0 {
			var ans = i
			if rocks[bags[i]] >= capacity[bags[i]] {
				ans++
			}
			return ans
		}
	}
	return len(bags)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maximumBags([]int{54, 18, 91, 49, 51, 45, 58, 54, 47, 91, 90, 20, 85, 20, 90, 49, 10, 84, 59, 29, 40, 9, 100, 1, 64, 71, 30, 46, 91},
		[]int{14, 13, 16, 44, 8, 20, 51, 15, 46, 76, 51, 20, 77, 13, 14, 35, 6, 34, 34, 13, 3, 8, 1, 1, 61, 5, 2, 15, 18}, 77))
}

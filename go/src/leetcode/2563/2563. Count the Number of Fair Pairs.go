package main

import (
	"sort"
)

/*
*
https://leetcode.cn/problems/count-the-number-of-fair-pairs/
lower - nums[j] <= nums[i] <= upper - nums[j]
*/
func countFairPairs(nums []int, lower int, upper int) (tot int64) {
	sort.Ints(nums)
	for i, x := range nums {
		l := sort.SearchInts(nums[:i], lower-x)
		r := sort.SearchInts(nums[:i], upper-x+1)
		tot += int64(r - l)
	}
	return
}

func main() {
	countFairPairs([]int{0, 1, 7, 4, 4, 5}, 3, 6)
	countFairPairs([]int{1, 7, 9, 2, 5}, 11, 11)

	//sorted := []int{0, 1, 4, 4, 7}
	//l := sort.Search(len(sorted), func(i int) bool {
	//	return sorted[i] > 4
	//})
	//fmt.Println(l)
}

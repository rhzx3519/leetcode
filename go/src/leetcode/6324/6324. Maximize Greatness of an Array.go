package main

import (
	"fmt"
	"sort"
)

/*
*
https://leetcode.cn/problems/maximize-greatness-of-an-array/
[1 1 1 2 3 3 5]

1,3,5,2,1,3,1
2,5,1,3,3,1,1
*/
func maximizeGreatness(nums []int) (tot int) {
	n := len(nums)
	sort.Ints(nums)
	perm := make([]int, n)
	copy(perm, nums)
	for len(perm) != 0 {
		x := perm[len(perm)-1]
		perm = perm[:len(perm)-1]
		j := sort.Search(len(nums), func(i int) bool {
			return nums[i] >= x
		})
		if j == 0 {
			break
		}
		copy(nums[j-1:], nums[j:])
		nums = nums[:len(nums)-1]
		tot++
	}
	return
}

func main() {
	fmt.Println(maximizeGreatness([]int{1, 3, 5, 2, 1, 3, 1}))
	fmt.Println(maximizeGreatness([]int{1, 2, 3, 4}))
}

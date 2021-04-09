package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	res = append(res, []int{})

	sort.Ints(nums)
	last := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		l, r := 0, len(res)
		if i > 0 && nums[i]==nums[i-1] {
			l = last
		}
		for j := l; j < r; j++	{
			t := []int{}
			t = append(t, res[j]...)
			t = append(t, nums[i])

			res = append(res, t)
		}
		last = r
	}

	return res
}

func main() {
	fmt.Println(subsetsWithDup([]int{1,2,2}))
}
package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	n := len(nums)
	perms := [][]int{}
	sort.Ints(nums)

	mapper := make(map[int]int)
	for _, num := range nums {
		mapper[num]++
	}

	var backtrace func(a []int)

	backtrace = func(a []int) {
		if len(a) == n {
			perms = append(perms, append([]int{}, a...))
			return
		}
		for i, num := range nums {
			if cnt, _ := mapper[num]; cnt <= 0 {
				continue
			}
			if i > 0 && nums[i-1] == nums[i] {
				continue
			}
			a = append(a, num)
			mapper[num]--
			backtrace(a)
			mapper[num]++
			a = a[:len(a)-1]
		}
	}

	backtrace([]int{})
	return perms
}

func main() {
	fmt.Println(permuteUnique([]int{1,1,2}))
	fmt.Println(permuteUnique([]int{1,3,2}))
}
package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	var ans = [][]int{}
	n := len(candidates)
	sort.Ints(candidates)

	var backtrace func(int, int, []int)
	backtrace = func(idx int, s int, arr []int) {
		if s <= 0|| idx == n {
			if s == 0 {
				ans = append(ans, append([]int{}, arr...))
			}
			return
		}

		for i := idx; i < n; i++ {
			if i > idx && candidates[i] == candidates[i - 1] {
				continue
			}
			arr = append(arr, candidates[i])
			backtrace(i + 1, s - candidates[i], arr)	// 只能选一次
			arr = arr[:len(arr) - 1]
		}
	}

	backtrace(0, target, []int{})
	return ans
}

func main() {
	fmt.Println(combinationSum2([]int{10,1,2,7,6,1,5}, 8))
}



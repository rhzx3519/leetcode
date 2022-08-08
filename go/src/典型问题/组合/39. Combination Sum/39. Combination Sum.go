package main

import "fmt"

/**
https://leetcode-cn.com/problems/combination-sum/
 */
func combinationSum(candidates []int, target int) [][]int {
	n := len(candidates)

	var ans = [][]int{}
	var backtrace func(int, int, []int)
	backtrace = func(idx int, s int, arr []int) {
		if s <= 0|| idx == n {
			if s == 0 {
				ans = append(ans, append([]int{}, arr...))
			}
			return
		}

		for i := idx; i < n; i++ {
			arr = append(arr, candidates[i])
			backtrace(i, s - candidates[i], arr)	// 可以选多次
			arr = arr[:len(arr) - 1]
		}
	}

	backtrace(0, target, []int{})
	fmt.Println(ans)
	return ans
}

func main() {
	combinationSum([]int{2,3,6,7}, 7)
}

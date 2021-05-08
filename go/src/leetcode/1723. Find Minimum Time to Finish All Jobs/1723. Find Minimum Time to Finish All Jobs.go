package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode-cn.com/problems/find-minimum-time-to-finish-all-jobs/
 */
func minimumTimeRequired(jobs []int, k int) int {
	s := 0
	sort.Sort(sort.Reverse(sort.IntSlice(jobs)))

	n := len(jobs)
	for i := 0; i < n; i++ {
		s += jobs[i]
	}
	l, r := jobs[0], s
	var mid int
	for l < r {
		mid = l + (r-l)>>1
		if check(jobs, k, mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func check(jobs []int, k int, limit int) bool {
	workloads := make([]int, k)
	var backtrace func(idx int) bool
	backtrace = func(idx int) bool {
		if idx == len(jobs)	{
			return true
		}
		for i := 0; i < k; i++ {
			if workloads[i] + jobs[idx] <= limit {
				workloads[i] += jobs[idx]
				if backtrace(idx+1) {
					return true
				}
				workloads[i] -= jobs[idx]
			}

			if workloads[i] == 0 || workloads[i] + jobs[idx] == limit {
				break
			}
		}
		return false
	}
	return backtrace(0)
}

func main() {
	fmt.Println(minimumTimeRequired([]int{3,2,3}, 3))
	fmt.Println(minimumTimeRequired([]int{8,7,4,1,2}, 2))
}


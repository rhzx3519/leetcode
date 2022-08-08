/**
@author ZhengHao Lou
Date    2022/06/13
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode.cn/problems/fair-distribution-of-cookies/
思路：二分 + 回溯查找,
和https://leetcode.cn/problems/find-minimum-time-to-finish-all-jobs/ 相同
*/
func distributeCookies(jobs []int, k int) int {
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
		if idx == len(jobs) {
			return true
		}
		for i := 0; i < k; i++ {
			if workloads[i]+jobs[idx] <= limit {
				workloads[i] += jobs[idx]
				if backtrace(idx + 1) {
					return true
				}
				workloads[i] -= jobs[idx]
			}
			
			if workloads[i] == 0 || workloads[i]+jobs[idx] == limit {
				break
			}
		}
		return false
	}
	return backtrace(0)
}

func main() {
	//fmt.Println(distributeCookies([]int{8, 15, 10, 20, 8}, 2))
	//fmt.Println(distributeCookies([]int{6, 1, 3, 2, 2, 4, 1, 2}, 3))
	//fmt.Println(distributeCookies([]int{10, 90}, 2))
	fmt.Println(distributeCookies([]int{11, 6, 14, 4}, 2))
}

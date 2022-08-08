/**
@author ZhengHao Lou
Date    2022/08/08
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/task-scheduler-ii/
*/
func taskSchedulerII(tasks []int, space int) int64 {
	n := len(tasks)
	last := make(map[int]int64)
	var day int64 = 0
	for i := 0; i < n; i++ {
		if r, ok := last[tasks[i]]; !ok || (day-r) > int64(space) {
			day++
		} else {
			day += int64(space) - (day - r - 1)
		}
		last[tasks[i]] = day
	}
	fmt.Println(last)
	fmt.Println(day)
	return day
}

func main() {
	taskSchedulerII([]int{1, 2, 1, 2, 3, 1}, 3)
	taskSchedulerII([]int{5, 8, 8, 5}, 2)
}

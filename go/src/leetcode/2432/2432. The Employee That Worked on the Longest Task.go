/**
@author ZhengHao Lou
Date    2022/10/10
*/
package main

/**
https://leetcode.cn/problems/the-employee-that-worked-on-the-longest-task/
*/
func hardestWorker(n int, logs [][]int) int {
	var j, mx = logs[0][0], logs[0][1]
	for i := 1; i < len(logs); i++ {
		t := logs[i][1] - logs[i-1][1]
		if t > mx {
			mx = t
			j = logs[i][0]
		} else if t == mx && logs[i][0] < j {
			j = logs[i][0]
		}
	}
	return j
}

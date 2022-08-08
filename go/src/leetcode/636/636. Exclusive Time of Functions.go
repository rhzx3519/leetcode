/**
@author ZhengHao Lou
Date    2022/08/07
*/
package main

import (
	"strconv"
	"strings"
)

func exclusiveTime(n int, logs []string) []int {
	var ans = make([]int, n)
	var st [][]int
	for _, log := range logs {
		strs := strings.Split(log, ":")
		id, _ := strconv.Atoi(strs[0])
		se := strs[1]
		t, _ := strconv.Atoi(strs[2])
		if se == "start" {
			st = append(st, []int{id, t})
		} else {
			x := st[len(st)-1]
			st = st[:len(st)-1]
			d := t - x[1] + 1
			ans[x[0]] += d
			if len(st) != 0 {
				ans[st[len(st)-1][0]] -= d
			}
		}
	}
	return ans
}

func main() {
	exclusiveTime(2, []string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"})
}

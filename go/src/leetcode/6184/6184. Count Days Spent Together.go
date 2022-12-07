/**
@author ZhengHao Lou
Date    2022/09/18
*/
package main

import (
	"strconv"
	"strings"
)

/**
https://leetcode.cn/problems/count-days-spent-together/
*/

var Month = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	var a1, a2 = decode(arriveAlice), decode(leaveAlice)
	var b1, b2 = decode(arriveBob), decode(leaveBob)
	if a1 > b1 {
		return countDaysTogether(arriveBob, leaveBob, arriveAlice, leaveAlice)
	}
	if b1 > a2 {
		return 0
	}
	return min(a2, b2) - max(a1, b1) + 1

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func decode(s string) int {
	strs := strings.Split(s, "-")
	mm, _ := strconv.Atoi(strs[0])
	dd, _ := strconv.Atoi(strs[1])
	var d int
	for i := 1; i < mm; i++ {
		d += Month[i-1]
	}
	d += dd
	return d
}

func main() {
	countDaysTogether("10-01", "10-31", "11-01", "12-31")
}

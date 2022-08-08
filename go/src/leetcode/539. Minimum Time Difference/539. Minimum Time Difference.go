/**
@author ZhengHao Lou
Date    2022/01/18
*/
package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

/**
https://leetcode-cn.com/problems/minimum-time-difference/
 */
func findMinDifference(timePoints []string) int {
	n := len(timePoints)
	times := []int{}
	for i := range timePoints {
		times = append(times, convert(timePoints[i]))
	}
	sort.Ints(times)
	fmt.Println(times)
	var d = math.MaxInt32>>1
	for i := range times {
		if i == 0 {
			d = min(d, times[0] + 24*60 - times[n-1])
		} else {
			d = min(d, times[i] - times[i-1])
		}
	}
	return d
}

func convert(s string) int {
	h, _ := strconv.Atoi(s[:2])
	m, _ := strconv.Atoi(s[3:])
	return h*60 + m
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	//fmt.Println(findMinDifference([]string{"23:59","00:00"}))
	//fmt.Println(findMinDifference([]string{"00:00","23:59","00:00"}))
	fmt.Println(findMinDifference([]string{"01:01","02:01"}))
}
/**
@author ZhengHao Lou
Date    2022/07/31
*/
package main

import (
	"fmt"
	"sort"
)

func maximumGroups(grades []int) int {
	sort.Ints(grades)
	n := len(grades)

	var arr [][]int
	var groups int
	var lastC, lastS int
	for i := 0; i < n; {
		var j = i
		var s, c int
		var tmp []int
		for ; j < n && (c <= lastC || s <= lastS); j++ {
			x := grades[j]
			c++
			s += x
			tmp = append(tmp, x)
		}
		arr = append(arr, tmp)
		if lastC >= c || lastS >= s {
			break
		}
		lastC = c
		lastS = s
		groups++
		i = j
	}

	fmt.Println(arr)
	return groups
}

func main() {
	fmt.Println(maximumGroups([]int{10, 6, 12, 7, 3, 5}))
	fmt.Println(maximumGroups([]int{10, 10}))
	fmt.Println(maximumGroups([]int{47, 2, 16, 17, 41, 4, 38, 23, 47}))
}

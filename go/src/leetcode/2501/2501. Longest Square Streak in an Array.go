/*
*
@author ZhengHao Lou
Date    2022/12/12
*/
package main

import (
	"fmt"
	"sort"
)

func longestSquareStreak(nums []int) (tot int) {
	tot = -1
	var counter = make(map[int]int)
	for _, x := range nums {
		counter[x]++
	}
	var a []int
	for i := range counter {
		a = append(a, i)
	}
	sort.Ints(a)
	n := len(a)
	f := make([]int, n)
	var index = make(map[int]int)
	for i := 0; i < n; i++ {
		f[i] = 1
		index[a[i]] = i
	}
	for i := range a {
		if f[i] > 1 {
			continue
		}
		var l int
		for j := i; counter[a[j]*a[j]] != 0; j = index[a[j]*a[j]] {
			l++
		}
		if l > 0 {
			tot = max(tot, l+1)
		}
	}

	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestSquareStreak([]int{4, 3, 6, 16, 8, 2}))
	fmt.Println(longestSquareStreak([]int{2, 3, 5, 6, 7}))
}

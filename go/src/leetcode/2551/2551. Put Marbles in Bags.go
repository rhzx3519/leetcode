package main

import (
	"fmt"
	"sort"
)

/*
*
https://leetcode.cn/problems/put-marbles-in-bags/
*/
func putMarbles(wt []int, k int) int64 {
	for i, x := range wt[1:] {
		wt[i] += x
	}
	wt = wt[:len(wt)-1]
	sort.Ints(wt)
	var ans int64
	for _, x := range wt[len(wt)-k+1:] {
		ans += int64(x)
	}
	for _, x := range wt[:k-1] {
		ans -= int64(x)
	}
	return ans
}

func main() {
	fmt.Println(putMarbles([]int{1, 3, 5, 1}, 2))
	fmt.Println(putMarbles([]int{1, 3}, 2))
}

package main

import (
	"fmt"
	"sort"
)

/*
*
https://leetcode.cn/problems/find-the-prefix-common-array-of-two-arrays/
*/
func findThePrefixCommonArray(A []int, B []int) []int {
	n := len(A)
	C, D := make([]int, n), make([]int, n)
	for i := range C {
		C[i] = i
		D[i] = i
	}
	sort.Slice(C, func(i, j int) bool {
		return A[C[i]] < A[C[j]]
	})
	sort.Slice(D, func(i, j int) bool {
		return B[D[i]] < B[D[j]]
	})

	diff := make([]int, n)
	for i := range A {
		diff[max(C[i], D[i])]++
	}

	var ans = make([]int, n)
	var c int
	for i := range diff {
		c += diff[i]
		ans[i] = c
	}

	fmt.Println(C, D, ans)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	findThePrefixCommonArray([]int{1, 3, 2, 4}, []int{3, 1, 2, 4})
	findThePrefixCommonArray([]int{2, 3, 1}, []int{3, 1, 2})
}

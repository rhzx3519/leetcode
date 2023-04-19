package main

import (
	"fmt"
)

// sum[i,j] = pre[j] - pre[i]
func example() {
	a := []int{0, 1, 0, 1, 1}
	n := len(a)
	pre := make([]int, n+1)
	for i := 1; i <= n; i++ {
		pre[i] = pre[i-1] + a[i-1]
	}
	fmt.Println(pre)
	fmt.Println(pre[n-1] - pre[0])
}

package main

import (
	"fmt"
	"math"
)

/*
*
https://leetcode.cn/problems/minimum-cost-tree-from-leaf-values/description/
*/
func mctFromLeafValues(arr []int) (tot int) {
	for len(arr) > 1 {
		var mi = math.MaxInt32
		var idx = 0
		for i := 0; i < len(arr)-1; i++ {
			if max(arr[i], arr[i+1]) < mi {
				mi = max(arr[i], arr[i+1])
				if arr[i] > arr[i+1] {
					idx = i + 1
				} else {
					idx = i
				}
			}
		}
		tot += mi * arr[idx]
		copy(arr[idx:], arr[idx+1:])
		arr = arr[:len(arr)-1]
		fmt.Println(arr)
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
	mctFromLeafValues([]int{6, 2, 4})
}

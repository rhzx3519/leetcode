/**
@author ZhengHao Lou
Date    2022/09/28
*/
package main

import (
	"fmt"
	"math"
)

/**
https://leetcode.cn/problems/get-kth-magic-number-lcci/

*/
func getKthMagicNumber(n int) int {
	var i, j, k int
	var que = []int{1}
	for tmp := 1; tmp < n; tmp++ {
		var cur = min(que[i]*3, que[j]*5, que[k]*7)
		que = append(que, cur)
		for ; que[i]*3 <= cur; i++ {
		}
		for ; que[j]*5 <= cur; j++ {
		}
		for ; que[k]*7 <= cur; k++ {
		}
	}
	fmt.Println(que)
	return que[len(que)-1]
}

func min(ints ...int) int {
	minVal := math.MaxInt32
	for _, i := range ints {
		if minVal > i {
			minVal = i
		}
	}
	return minVal
}

func main() {
	getKthMagicNumber(8)
}

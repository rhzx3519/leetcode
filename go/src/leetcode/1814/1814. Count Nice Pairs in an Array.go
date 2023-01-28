/**
@author ZhengHao Lou
Date    2023/01/17
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/count-nice-pairs-in-an-array/
*/
const mod int = 1e9 + 7

func countNicePairs(nums []int) (tot int) {
	var rev = func(x int) int {
		var y int
		for ; x != 0; x /= 10 {
			y = y*10 + x%10
		}
		return y
	}
	var counter = make(map[int]int)
	for _, x := range nums {
		k := x - rev(x)
		counter[k]++
		tot = (tot + counter[k] - 1) % mod
	}
	return
}

func main() {
	fmt.Println(countNicePairs([]int{42, 11, 1, 97}))
	fmt.Println(countNicePairs([]int{13, 10, 35, 24, 76}))
}

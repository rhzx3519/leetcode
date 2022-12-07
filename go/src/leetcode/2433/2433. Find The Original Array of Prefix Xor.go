/**
@author ZhengHao Lou
Date    2022/10/10
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/find-the-original-array-of-prefix-xor/
*/
func findArray(pref []int) []int {
	n := len(pref)
	arr := make([]int, n)
	arr[0] = pref[0]
	for i := 1; i < n; i++ {
		arr[i] = pref[i] ^ pref[i-1]
	}

	fmt.Println(arr)
	return arr
}

func main() {
	findArray([]int{5, 2, 0, 3, 1})
}

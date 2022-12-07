/**
@author ZhengHao Lou
Date    2022/08/25
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/find-k-closest-elements/
*/
func findClosestElements(arr []int, k int, x int) []int {
	l, r := 0, len(arr)-1
	for r-l >= k {
		if abs(arr[l]-x) <= abs(arr[r]-x) {
			r--
		} else {
			l++
		}
	}
	fmt.Println(arr[l : r+1])
	return arr[l : r+1]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3)
	findClosestElements([]int{1, 2, 3, 4, 5}, 4, -1)
}

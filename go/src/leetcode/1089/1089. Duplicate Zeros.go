/**
@author ZhengHao Lou
Date    2022/06/17
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/duplicate-zeros/
*/
func duplicateZeros(arr []int) {
	n := len(arr)
	for i := 0; i < n; {
		if arr[i] == 0 {
			if i+1 < n {
				if i+2 < n {
					copy(arr[i+2:], arr[i+1:])
				}
				arr[i+1] = 0
			}
			i += 2
		} else {
			i++
		}
	}
	fmt.Println(arr)
}

func main() {
	duplicateZeros([]int{1, 0, 2, 3, 0, 4, 5, 0})
}

/**
@author ZhengHao Lou
Date    2022/02/20
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/1-bit-and-2-bit-characters/
*/
func isOneBitCharacter(bits []int) bool {
	n := len(bits)
	var i int
	for i < n-1 {
		if bits[i] == 0 {
			i++
		} else {
			i += 2
		}
	}
	return i == n-1
}

func main() {
	fmt.Println(isOneBitCharacter([]int{1, 0, 0}))
	fmt.Println(isOneBitCharacter([]int{1, 1, 1, 0}))
}

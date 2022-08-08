/**
@author ZhengHao Lou
Date    2022/03/13
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/utf-8-validation/
*/
func validUtf8(data []int) bool {
	n := len(data)
	var i int
	for i < n {
		if data[i]>>7&1 == 0 {
			i++
		} else if data[i]>>5 == 1<<3-2 {
			var j = 1
			for ; i+j < n && j < 2; j++ {
				if data[i+j]>>6 != 2 {
					break
				}
			}
			if j != 2 {
				return false
			}
			i = i + j
		} else if data[i]>>4 == 1<<4-2 {
			var j = 1
			for ; i+j < n && j < 3; j++ {
				if data[i+j]>>6 != 2 {
					break
				}
			}
			if j != 3 {
				return false
			}
			i = i + j
		} else if data[i]>>3 == 1<<5-2 {
			var j = 1
			for ; i+j < n && j < 4; j++ {
				if data[i+j]>>6 != 2 {
					break
				}
			}
			if j != 4 {
				return false
			}
			i = i + j
		} else {
			return false
		}
	}
	return true
}

func main() {
	//fmt.Println(validUtf8([]int{197, 130, 1}))
	//fmt.Println(validUtf8([]int{235, 140, 4}))
	fmt.Println(validUtf8([]int{240, 162, 138, 147}))
}

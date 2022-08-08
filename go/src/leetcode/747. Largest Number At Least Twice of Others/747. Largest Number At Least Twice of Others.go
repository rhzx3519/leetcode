/**
@author ZhengHao Lou
Date    2022/01/13
*/
package main

import "fmt"

func dominantIndex(nums []int) int {
	var idx int
	var a, b = -1, -1
	for i, num := range nums {
		if num >= a {
			b = a
			a = num
			idx = i
		} else if num >= b {
			b = num
		}

	}
	fmt.Println(a, b)
	if a < b<<1 {
		return -1
	}
	return idx
}

func main() {
	fmt.Println(dominantIndex([]int{3,6,1,0}))
	//fmt.Println(dominantIndex([]int{1,2,3,4}))
	//fmt.Println(dominantIndex([]int{1}))
	fmt.Println(dominantIndex([]int{0,0,3,2}))
}

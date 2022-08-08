/**
@author ZhengHao Lou
Date    2022/03/15
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/count-number-of-maximum-bitwise-or-subsets/
*/
func countMaxOrSubsets(nums []int) int {
	var mx, c int
	subsets := []int{0}
	for _, num := range nums {
		tmp := []int{}
		for _, sub := range subsets {
			x := sub | num
			if x > mx {
				mx = x
				c = 1
			} else if x == mx {
				c++
			}
			tmp = append(tmp, sub|num)
		}
		subsets = append(subsets, tmp...)
	}

	return c
}

func main() {
	fmt.Println(countMaxOrSubsets([]int{3, 1}))
	fmt.Println(countMaxOrSubsets([]int{2, 2, 2}))
	fmt.Println(countMaxOrSubsets([]int{3, 2, 1, 5}))
}

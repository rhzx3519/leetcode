/*
*
@author ZhengHao Lou
Date    2022/12/16
*/
package main

import "fmt"

/*
*
https://leetcode.cn/problems/minimum-elements-to-add-to-form-a-given-sum/
*/
func minElements(nums []int, limit int, goal int) (tot int) {
	var s int
	for i := range nums {
		s += nums[i]
	}
	x := goal - s
	if x == 0 {
		return 0
	}
	if x < 0 {
		x = -x
	}
	tot += x / limit
	if x%limit != 0 {
		tot++
	}
	return
}

func main() {
	fmt.Println(minElements([]int{1, -1, 1}, 3, -4))
	fmt.Println(minElements([]int{1, -10, 9, 1}, 100, 0))
}

/**
@author ZhengHao Lou
Date    2021/10/18
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/count-number-of-maximum-bitwise-or-subsets/
 */
func countMaxOrSubsets(nums []int) int {
	subsets := []int{0}
	for _, num := range nums {
		tmp := []int{}
		for _, sub := range subsets {
			tmp = append(tmp, sub | num)
		}
		subsets = append(subsets, tmp...)
	}
	var count, max int
	for _, sub := range subsets {
		if sub > max {
			max= sub
			count = 1
		} else if sub == max {
			count++
		}
	}
	fmt.Println(count)
	return count
}

func main() {
	countMaxOrSubsets([]int{3,1})
	countMaxOrSubsets([]int{2,2,2})
	countMaxOrSubsets([]int{3,2,1,5})
}
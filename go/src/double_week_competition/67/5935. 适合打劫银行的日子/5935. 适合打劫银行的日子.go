/**
@author ZhengHao Lou
Date    2021/12/11
*/
package main

import "fmt"

/**
https://leetcode-cn.com/contest/biweekly-contest-67/problems/find-good-days-to-rob-the-bank/
 */
func goodDaysToRobBank(security []int, time int) []int {
	n := len(security)
	var c int
	var left, right = make([]int, n), make([]int, n)
	for i := 1; i < n; i++ {
		if security[i] <= security[i-1] {
			c++
		} else {
			c = 0
		}
		left[i]	= c
	}

	c = 0
	for i := n-2; i >= 0; i-- {
		if security[i] <= security[i+1] {
			c++
		} else {
			c = 0
		}
		right[i] = c
	}

	var ans = []int{}
	for i := 0; i < n; i++ {
		if left[i] >= time && right[i] >= time {
			ans = append(ans, i)
		}
	}
	//fmt.Println(left)
	//fmt.Println(right)
	fmt.Println(ans)
	return ans
}

func main() {
	goodDaysToRobBank([]int{5,3,3,3,5,6,2}, 2)
	goodDaysToRobBank([]int{1,1,1,1,1}, 0)
	goodDaysToRobBank([]int{1,2,3,4,5,6}, 2)
	goodDaysToRobBank([]int{1}, 5)
	goodDaysToRobBank([]int{1,2,3,4}, 1)
}

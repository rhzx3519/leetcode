/**
@author ZhengHao Lou
Date    2022/03/06
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/find-good-days-to-rob-the-bank/
*/
func goodDaysToRobBank(security []int, time int) []int {
	n := len(security)
	ascend, descend := make([]int, n), make([]int, n)
	var c = 1
	for i := 1; i < n; i++ {
		if security[i-1] >= security[i] {
			descend[i] = c
			c++
		} else {
			descend[i] = 0
			c = 1
		}
	}
	c = 1
	for i := n - 2; i >= 0; i-- {
		if security[i+1] >= security[i] {
			ascend[i] = c
			c++
		} else {
			ascend[i] = 0
			c = 1
		}
	}

	var rob []int
	for i := 0; i < n; i++ {
		if ascend[i] >= time && descend[i] >= time {
			rob = append(rob, i)
		}
	}
	//fmt.Println(descend)
	//fmt.Println(ascend)
	fmt.Println(rob)
	return rob
}

func main() {
	goodDaysToRobBank([]int{5, 3, 3, 3, 5, 6, 2}, 2)
	goodDaysToRobBank([]int{1, 1, 1, 1, 1}, 0)
	goodDaysToRobBank([]int{1, 2, 3, 4, 5, 6}, 2)
	goodDaysToRobBank([]int{1}, 5)
}

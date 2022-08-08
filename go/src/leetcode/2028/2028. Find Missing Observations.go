/**
@author ZhengHao Lou
Date    2022/03/27
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/find-missing-observations/
*/
func missingRolls(rolls []int, mean int, n int) []int {
	m := len(rolls)
	var sm int
	for i := range rolls {
		sm += rolls[i]
	}
	sn := mean*(m+n) - sm
	if sn < n || sn > 6*n {
		return []int{}
	}

	x := sn / n
	r := sn % n
	var miss []int
	for i := 0; i < n; i++ {
		miss = append(miss, x)
		if r > 0 {
			t := min(r, 6-x)
			r -= t
			miss[i] += t
		}
	}

	return miss
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(missingRolls([]int{1, 2, 3, 4}, 6, 4))
	//fmt.Println(missingRolls([]int{3, 2, 4, 3}, 4, 2))
	//fmt.Println(missingRolls([]int{1, 5, 6}, 3, 4))
	//fmt.Println(missingRolls([]int{1}, 3, 1))
}

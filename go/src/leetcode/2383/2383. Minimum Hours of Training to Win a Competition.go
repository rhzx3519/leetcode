/**
@author ZhengHao Lou
Date    2022/08/22
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/minimum-hours-of-training-to-win-a-competition/
*/
func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	n := len(energy)
	var eng, exp int = initialEnergy, initialExperience
	var m1, m2 int
	for i := 0; i < n; i++ {
		if eng <= energy[i] {
			m1 = max(m1, energy[i]-eng+1)
		}
		if exp <= experience[i] {
			m2 = max(m2, experience[i]-exp+1)
		}

		eng -= energy[i]
		exp += experience[i]
	}
	fmt.Println(m1, m2)
	return m1 + m2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//minNumberOfHours(5, 3, []int{1, 4, 3, 2}, []int{2, 6, 3, 1})
	//minNumberOfHours(2, 4, []int{1}, []int{3})
	// 2
	minNumberOfHours(5, 3, []int{1, 4}, []int{2, 5})
}

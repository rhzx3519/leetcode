/**
@author ZhengHao Lou
Date    2022/04/18
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/minimum-rounds-to-complete-all-tasks/
*/
func minimumRounds(tasks []int) int {
	counter := make(map[int]int)
	for _, t := range tasks {
		counter[t]++
	}
	var round int
	for _, c := range counter {
		x := can(c)
		if x == -1 {
			return -1
		} else {
			round += x
		}
	}
	fmt.Println(counter, round)
	return round
}

func can(x int) int {
	for i := 0; i<<1 <= x; i++ {
		if (x-i<<1)%3 == 0 {
			return (x-i<<1)/3 + i
		}
	}
	return -1
}

func main() {
	minimumRounds([]int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4})
	minimumRounds([]int{2, 3, 3})
}

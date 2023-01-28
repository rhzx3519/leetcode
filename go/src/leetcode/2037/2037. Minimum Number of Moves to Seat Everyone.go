/*
*
@author ZhengHao Lou
Date    2022/12/31
*/
package main

import (
	"fmt"
	"sort"
)

/*
*
https://leetcode.cn/problems/minimum-number-of-moves-to-seat-everyone/
*/
func minMovesToSeat(seats []int, students []int) (tot int) {
	sort.Ints(seats)
	sort.Ints(students)
	for i := range seats {
		tot += abs(students[i] - seats[i])
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	//fmt.Println(minMovesToSeat([]int{3, 1, 5}, []int{2, 7, 4}))
	//fmt.Println(minMovesToSeat([]int{4, 1, 5, 9}, []int{1, 3, 2, 6}))
	fmt.Println(minMovesToSeat([]int{2, 2, 6, 6}, []int{1, 3, 2, 6}))
}

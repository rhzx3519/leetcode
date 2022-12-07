/*
*
@author ZhengHao Lou
Date    2022/12/05
*/
package main

import "fmt"

/*
*
https://leetcode.cn/problems/divide-players-into-teams-of-equal-skill/
*/
func dividePlayers(skill []int) (tot int64) {
	var s int
	var counter = make(map[int]int)
	for i := range skill {
		s += skill[i]
		counter[skill[i]]++
	}
	n := len(skill)
	if s%(n>>1) != 0 {
		return -1
	}
	ts := s / (n >> 1)

	for _, x := range skill {
		if counter[x] == 0 {
			continue
		}
		if counter[ts-x] == 0 {
			return -1
		}
		tot += int64(x * (ts - x))
		counter[ts-x]--
		counter[x]--
	}
	return
}

func main() {
	fmt.Println(dividePlayers([]int{3, 2, 5, 1, 3, 4}))
	fmt.Println(dividePlayers([]int{3, 4}))
	fmt.Println(dividePlayers([]int{1, 1, 2, 3}))
}

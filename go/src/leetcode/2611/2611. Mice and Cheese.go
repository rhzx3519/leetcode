package main

import (
	"fmt"
	"sort"
)

/*
*
https://leetcode.cn/problems/mice-and-cheese/
*/
func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	n := len(reward1)
	var ids = make([]int, n)
	for i := range reward1 {
		ids[i] = i
	}
	sort.Slice(ids, func(i, j int) bool {
		return reward1[ids[i]]-reward2[ids[i]] > reward1[ids[j]]-reward2[ids[j]]
	})

	var tot int
	for i := 0; i < n; i++ {
		if i < k {
			tot += reward1[ids[i]]
		} else {
			tot += reward2[ids[i]]
		}
	}

	return tot
}

func main() {
	//fmt.Println(miceAndCheese([]int{3, 4, 1, 1}, []int{4, 4, 1, 1}, 2))
	//fmt.Println(miceAndCheese([]int{1, 1, 3, 4}, []int{4, 4, 1, 1}, 2))
	//fmt.Println(miceAndCheese([]int{1, 1}, []int{1, 1}, 2))
	fmt.Println(miceAndCheese([]int{1, 4, 4, 6, 4}, []int{6, 5, 3, 6, 1}, 1)) // 24
}

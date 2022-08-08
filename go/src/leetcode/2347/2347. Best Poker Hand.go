/**
@author ZhengHao Lou
Date    2022/07/25
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/best-poker-hand/
*/
func bestHand(ranks []int, suits []byte) string {
	var c = true
	var lastS byte = '0'
	for i := range suits {
		if lastS == '0' {
			lastS = suits[i]
		} else if lastS != suits[i] {
			c = false
		}
	}
	if c {
		return "Flush"
	}
	
	var rankMap = map[int]int{}
	for i := range ranks {
		rankMap[ranks[i]]++
	}
	
	var mx int
	for i := range rankMap {
		if rankMap[i] >= mx {
			mx = rankMap[i]
		}
	}
	
	if mx >= 3 {
		return "Three of a Kind"
	} else if mx >= 2 {
		return "Pair"
	}
	
	return "High Card"
}

func main() {
	fmt.Println(bestHand([]int{13, 2, 3, 1, 9}, []byte{'a', 'a', 'a', 'a', 'a'}))
	fmt.Println(bestHand([]int{13, 2, 3, 1, 9}, []byte{'a', 'a', 'a', 'a', 'a'}))
}

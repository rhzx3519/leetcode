/**
@author ZhengHao Lou
Date    2021/12/02
*/
package main

import (
	"fmt"
	"sort"
	"strconv"
)

/**
https://leetcode-cn.com/problems/relative-ranks/
 */
func findRelativeRanks(score []int) []string {
	ranks := []int{}
	n := len(score)
	for i := range score {
		ranks = append(ranks, i)
	}
	sort.Slice(ranks, func(i, j int) bool {
		return score[ranks[i]] > score[ranks[j]]
	})
	ans := make([]string, n)
	for i := range ranks {
		var s string
		switch i {
		case 0:
			s = "Gold Medal"
		case 1:
			s = "Silver Medal"
		case 2:
			s = "Bronze Medal"
		default:
			s = strconv.Itoa(i + 1)
		}
		ans[ranks[i]] = s
	}
	return ans
}

func main() {
	fmt.Println(findRelativeRanks([]int{5,4,3,2,1}))
	fmt.Println(findRelativeRanks([]int{10,3,8,9,4}))
}

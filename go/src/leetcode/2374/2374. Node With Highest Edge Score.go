/**
@author ZhengHao Lou
Date    2022/08/15
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/node-with-highest-edge-score/
*/
func edgeScore(edges []int) int {
	n := len(edges)
	scores := make([]int, n)
	for i, j := range edges {
		scores[j] += i
	}
	var ans, mx = n, 0
	for i := range scores {
		if scores[i] > mx {
			mx = scores[i]
			ans = i
		} else if scores[i] == mx && i < ans {
			ans = i
		}
	}
	fmt.Print(scores)
	return ans
}

func main() {
	edgeScore([]int{1, 0, 0, 0, 0, 7, 7, 5})
}

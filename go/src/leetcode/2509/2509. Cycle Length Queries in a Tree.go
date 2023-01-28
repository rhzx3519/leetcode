/*
*
@author ZhengHao Lou
Date    2022/12/19
*/
package main

import (
	"fmt"
)

/*
*
https://leetcode.cn/problems/cycle-length-queries-in-a-tree/
1<<0
2*n, 2*n+1
*/
func cycleLengthQueries(n int, queries [][]int) []int {
	var ans = make([]int, len(queries))
	for i := range queries {
		a, b := queries[i][0], queries[i][1]
		var c int
		for a != b {
			if a > b {
				a >>= 1
				c++
			} else {
				b >>= 1
				c++
			}
		}
		ans[i] = c + 1
	}
	return ans
}

func main() {
	fmt.Println(cycleLengthQueries(3, [][]int{{5, 3}, {4, 7}, {2, 3}}))
	fmt.Println(cycleLengthQueries(2, [][]int{{1, 2}}))
	fmt.Println(cycleLengthQueries(5, [][]int{{17, 21}, {23, 5}}))
}

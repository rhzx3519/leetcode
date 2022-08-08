/**
@author ZhengHao Lou
Date    2021/11/01
*/
package main

import (
	"fmt"
)

/**
https://leetcode-cn.com/problems/distribute-candies/
 */
func distributeCandies(candyType []int) int {
	n := len(candyType)
	var t int
	count := make(map[int]int)
	for _, c := range candyType {
		if count[c] == 0 {
			t++
			if t >= n>>1 {
				return t
			}
		}
		count[c]++
	}
	return len(count)
}

func main() {
	fmt.Println(distributeCandies([]int{1,1,2,2,3,3}))
	fmt.Println(distributeCandies([]int{1,1,2,3}))
}
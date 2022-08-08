/**
@author ZhengHao Lou
Date    2021/11/22
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/two-furthest-houses-with-different-colors/
 */
func maxDistance(colors []int) int {
	var ans int
	n := len(colors)
	l, r := 0, n-1
	for l < r && colors[l] == colors[r] {
		l++
	}
	ans = max(ans, r-l)
	l, r = 0, n-1
	for l < r && colors[l] == colors[r] {
		r--
	}
	ans = max(ans, r-l)

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxDistance([]int{1,1,1,6,1,1,1}))
	fmt.Println(maxDistance([]int{1,8,3,8,3}))
	fmt.Println(maxDistance([]int{0,1}))
	fmt.Println(maxDistance([]int{9,9,9,18,9,9,9,9,9,18}))	// 9
	fmt.Println(maxDistance([]int{6,6,6,6,6,6,6,6,6,19,19,6,6}))
}
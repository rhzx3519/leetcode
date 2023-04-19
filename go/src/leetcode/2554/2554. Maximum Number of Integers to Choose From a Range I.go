package main

import "fmt"

/*
*
https://leetcode.cn/problems/maximum-number-of-integers-to-choose-from-a-range-i/description/
*/
func maxCount(banned []int, n, maxSum int) (ans int) {
	has := map[int]bool{}
	for _, v := range banned {
		has[v] = true
	}
	for i := 1; i <= n && i <= maxSum; i++ {
		if !has[i] {
			fmt.Println(i, maxSum)
			maxSum -= i
			ans++
		}
	}
	return
}

func main() {
	maxCount([]int{1, 6, 5}, 5, 6)
}

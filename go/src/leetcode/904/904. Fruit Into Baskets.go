/**
@author ZhengHao Lou
Date    2022/10/17
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/fruit-into-baskets/
*/
func totalFruit(fruits []int) (c int) {
	var counter = make(map[int]int)
	var l int
	for r := range fruits {
		counter[fruits[r]]++
		for ; l < r && len(counter) > 2; l++ {
			counter[fruits[l]]--
			if counter[fruits[l]] == 0 {
				delete(counter, fruits[l])
			}
		}
		c = max(c, r-l+1)
	}
	fmt.Println(c)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	totalFruit([]int{1, 2, 1})
	totalFruit([]int{0, 1, 2, 2})
	totalFruit([]int{1, 2, 3, 2, 2})
	totalFruit([]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4})
}

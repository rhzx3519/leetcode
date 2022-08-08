/**
@author ZhengHao Lou
Date    2022/01/24
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/find-all-lonely-numbers-in-the-array/
*/
func findLonely(nums []int) []int {
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}
	ans := []int{}
	for x, c := range counter {
		if c == 1 && counter[x-1] == 0 && counter[x+1] == 0 {
			ans = append(ans, x)
		}
	}
	fmt.Println(counter)
	return ans
}

func main() {
	findLonely([]int{10, 6, 5, 8})
}

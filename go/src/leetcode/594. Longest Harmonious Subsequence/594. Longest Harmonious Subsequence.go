/**
@author ZhengHao Lou
Date    2021/11/20
*/
package main

import "fmt"

func findLHS(nums []int) int {
	mapper := make(map[int]int)
	for _, num := range nums {
		mapper[num]++
	}

	var ans int
	for _, num := range nums {
		if mapper[num-1] != 0 {
			ans = max(ans, mapper[num] + mapper[num-1])
		}
		if mapper[num+1] != 0 {
			ans = max(ans, mapper[num] + mapper[num+1])
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(findLHS([]int{1,3,2,2,5,2,3,7}))
	fmt.Println(findLHS([]int{1,2,3,4}))
	fmt.Println(findLHS([]int{1,1,1,1}))
}
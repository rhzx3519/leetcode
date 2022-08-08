/**
@author ZhengHao Lou
Date    2021/11/22
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/k-diff-pairs-in-an-array/
 */
func findPairs(nums []int, k int) int {
	var set = make(map[[2]int]int)
	var vis = make(map[int]int)
	for _, num := range nums {
		if _, ok := vis[num - k]; ok {
			set[[2]int{num - k, num}]++
		}
		if _, ok := vis[num + k]; ok {
			set[[2]int{num, num + k}]++
		}
		vis[num]++
	}

	fmt.Println(len(set))
	return len(set)
}

func main() {
	findPairs([]int{3,1,4,1,5}, 2)
	findPairs([]int{1,2,3,4,5}, 1)
	findPairs([]int{1,3,1,5,4}, 0)
	findPairs([]int{1,2,4,4,3,3,0,9,2,3}, 3)
	findPairs([]int{-1,-2,-3}, 1)
}




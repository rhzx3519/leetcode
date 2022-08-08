/**
@author ZhengHao Lou
Date    2021/12/02
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode-cn.com/problems/divide-array-in-sets-of-k-consecutive-numbers/
 */
func isPossibleDivide(nums []int, k int) bool {
	n := len(nums)
	if n%k != 0 {
		return false
	}
	if k == 1 {
		return true
	}
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}
	sort.Ints(nums)
	for _, num := range nums {
		if counter[num] == 0 {
			continue
		}
		for i := 0; i < k; i++ {
			if counter[num + i] == 0 {
				return false
			}
			counter[num + i]--
		}
	}

	return true
}

func main() {
	fmt.Println(isPossibleDivide([]int{1,2,3,3,4,4,5,6}, 4))
	fmt.Println(isPossibleDivide([]int{3,2,1,2,3,4,3,4,5,9,10,11}, 3))
	fmt.Println(isPossibleDivide([]int{3,3,2,2,1,1}, 3))
	fmt.Println(isPossibleDivide([]int{1,2,3,4}, 3))
}
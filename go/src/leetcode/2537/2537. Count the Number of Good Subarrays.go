/**
@author ZhengHao Lou
Date    2023/01/16
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/count-the-number-of-good-subarrays/
思路：滑动窗口
*/
func countGood(nums []int, k int) (tot int64) {
	counter := make(map[int]int)
	var pairs int
	var l int
	for _, x := range nums {
		pairs += counter[x]
		counter[x]++
		for pairs-counter[nums[l]]+1 >= k {
			counter[nums[l]]--
			pairs -= counter[nums[l]]
			l++
		}
		if pairs >= k {
			tot += int64(l + 1)
		}
	}
	return
}

func main() {
	fmt.Println(countGood([]int{1, 1, 1, 1, 1}, 10))
	fmt.Println(countGood([]int{3, 1, 4, 3, 2, 2, 4}, 2))
	
	fmt.Println(countGood([]int{2, 3, 1, 3, 2, 3, 3, 3, 1, 1, 3, 2, 2, 2}, 18)) // 9
}

/**
@author ZhengHao Lou
Date    2022/08/08
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/count-number-of-bad-pairs/
j - i != nums[j] - nums[i]
nums[i] - i != nums[j] - j
*/
func countBadPairs(nums []int) int64 {
	var ans int64
	counter := make(map[int]int)
	for j := range nums {
		ans += int64((j + 0) - counter[nums[j]-j])
		counter[nums[j]-j]++
	}
	fmt.Println(ans)
	return ans
}

func main() {
	countBadPairs([]int{4, 1, 3, 3})
	countBadPairs([]int{1, 2, 3, 4, 5})
}

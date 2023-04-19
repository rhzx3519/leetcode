package main

import "fmt"

/*
*
https://leetcode.cn/problems/count-increasing-quadruplets/
i < j < k < l
nums[i] < nums[k] < nums[j] < nums[l]
*/
func countQuadruplets(nums []int) int64 {
	var tot int64
	var cnt = make(map[int]int)
	for i := range nums {
		var large int
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				// cnt[j]表示类似132三元组的数量
				tot += int64(cnt[j])
				large++
			} else {
				// nums[i] < nums[j]
				// i < j < k, nums[i] < nums[k] < nums[j]
				cnt[j] += large // 统计类似132，j为2的三元组的数量
			}
		}
	}
	return tot
}

func main() {
	fmt.Println(countQuadruplets([]int{1, 3, 2, 4, 5}))
	fmt.Println(countQuadruplets([]int{1, 2, 3, 4}))
}

package main

import "sort"

/**
https://leetcode.cn/problems/minimum-number-game/?envType=daily-question&envId=2024-07-12
*/
func numberGame(nums []int) []int {
    sort.Ints(nums)
    for i := 0; i < len(nums); i += 2 {
        nums[i], nums[i+1] = nums[i+1], nums[i]
    }
    return nums
}

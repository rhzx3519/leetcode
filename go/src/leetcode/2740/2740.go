package main

import (
    "math"
    "sort"
)

/**
https://leetcode.cn/problems/find-the-value-of-the-partition/?envType=daily-question&envId=2024-07-26
*/
func findValueOfPartition(nums []int) int {
    sort.Ints(nums)
    var ans = math.MaxInt32
    for i := range nums {
        if i > 0 {
            ans = min(ans, min(nums[i]-nums[i-1]))
        }
    }
    return ans
}

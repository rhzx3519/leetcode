package main

import (
    "math/bits"
    "slices"
)

/**
https://leetcode.cn/problems/find-if-array-can-be-sorted/?envType=daily-question&envId=2024-07-13
*/
func canSortArray(nums []int) bool {
    for i, n := 0, len(nums); i < n; {
        start := i
        ones := bits.OnesCount(uint(nums[i]))
        i++
        for i < n && bits.OnesCount(uint(nums[i])) == ones {
            i++
        }
        slices.Sort(nums[start:i])
    }
    return slices.IsSorted(nums)
}

package main

import (
    "slices"
    "sort"
)

/**
https://leetcode.cn/problems/minimum-number-of-operations-to-make-array-continuous/description/?envType=daily-question&envId=2024-04-08
思路：滑动窗口
*/
func minOperations(nums []int) int {
    n := len(nums)
    sort.Ints(nums)
    a := slices.Compact(nums)
    var ans, left int
    for i, x := range a {
        for a[left] < x-n+1 {
            left++
        }
        ans = max(ans, i-left+1)
    }
    return n - ans
}

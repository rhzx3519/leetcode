package main

/**
https://leetcode.cn/problems/find-pivot-index/?envType=daily-question&envId=2024-07-08
*/
func pivotIndex(nums []int) int {
    var s int
    for i := range nums {
        s += nums[i]
    }
    var c int
    for i := range nums {
        if c == s-nums[i] {
            return i
        }
        c += nums[i]
        s -= nums[i]
    }
    return -1
}

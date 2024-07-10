package main

/**
https://leetcode.cn/problems/count-the-number-of-incremovable-subarrays-i/?envType=daily-question&envId=2024-07-10
*/
func incremovableSubarrayCount(nums []int) int {
    n := len(nums)
    res := 0
    for i := 0; i < n; i++ {
        for j := i; j < n; j++ {
            if isIncreasing(nums, i, j) {
                res++
            }
        }
    }
    return res
}

func isIncreasing(nums []int, l, r int) bool {
    for i := 1; i < len(nums); i++ {
        if i >= l && i <= r+1 {
            continue
        }
        if nums[i] <= nums[i-1] {
            return false
        }
    }
    if l-1 >= 0 && r+1 < len(nums) && nums[r+1] <= nums[l-1] {
        return false
    }
    return true
}

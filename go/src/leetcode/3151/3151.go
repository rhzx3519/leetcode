package main

/**
https://leetcode.cn/problems/special-array-i/?envType=daily-question&envId=2024-08-13
*/
func isArraySpecial(nums []int) bool {
    for i := 0; i < len(nums)-1; i++ {
        if nums[i]%2 == nums[i+1]%2 {
            return false
        }
    }
    return true
}

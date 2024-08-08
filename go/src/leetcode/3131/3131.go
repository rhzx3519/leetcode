package main

/**
https://leetcode.cn/problems/find-the-integer-added-to-array-i/?envType=daily-question&envId=2024-08-08
*/
func addedInteger(nums1 []int, nums2 []int) int {
    var x int
    for i := range nums1 {
        x += nums2[i] - nums1[i]
    }
    return x / len(nums1)
}

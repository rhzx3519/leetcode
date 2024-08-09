package main

import "sort"

/**
https://leetcode.cn/problems/find-the-integer-added-to-array-ii/?envType=daily-question&envId=2024-08-09
*/
func minimumAddedInteger(nums1 []int, nums2 []int) int {
    sort.Ints(nums1)
    sort.Ints(nums2)
    for i := 2; i >= 0; i-- {
        left, right := i+1, 1
        for left < len(nums1) && right < len(nums2) {
            if nums1[left]-nums2[right] == nums1[i]-nums2[0] {
                right++
            }
            left++
        }
        if right == len(nums2) {
            return nums2[0] - nums1[i]
        }
    }
    return 0
}

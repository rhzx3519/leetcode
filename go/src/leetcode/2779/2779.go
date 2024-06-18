package main

import "sort"

/**
https://leetcode.cn/problems/maximum-beauty-of-an-array-after-applying-operation/?envType=daily-question&envId=2024-06-15
*/
func maximumBeauty(nums []int, k int) (tot int) {
    sort.Ints(nums)
    var l int
    for r := range nums {
        for nums[r]-nums[l] > k*2 {
            l++
        }
        tot = max(tot, r-l+1)
    }

    return
}

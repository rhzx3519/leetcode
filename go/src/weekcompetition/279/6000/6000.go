/**
@author ZhengHao Lou
Date    2022/02/06
*/
package main

import "sort"

func sortEvenOdd(nums []int) []int {
    var odd, even []int
    for i := range nums {
        if i&1 == 1 {
            odd = append(odd, nums[i])
        } else {
            even = append(even, nums[i])
        }
    }
    sort.Ints(even)
    sort.Slice(odd, func(i, j int) bool {
        return odd[i] >= odd[j]
    })
    var ans []int
    for i, j := 0, 0; i < len(even) || j < len(odd); i, j = i+1, j+1 {
        if i < len(even) {
            ans = append(ans, even[i])
        }
        if j < len(odd) {
            ans = append(ans, odd[j])
        }
    }
    return ans
}

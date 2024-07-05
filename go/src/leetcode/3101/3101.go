package main

import "fmt"

/**
https://leetcode.cn/problems/count-alternating-subarrays/?envType=daily-question&envId=2024-07-06
*/
func countAlternatingSubarrays(nums []int) int64 {
    var c int
    var tot int64
    for i := range nums {
        if i > 0 && nums[i] == nums[i-1] {
            c = 1
        } else {
            c++
        }
        tot += int64(c)
    }
    return tot
}

func main() {
    fmt.Println(countAlternatingSubarrays([]int{0, 1, 1, 1}))
    fmt.Println(countAlternatingSubarrays([]int{1, 0, 1, 0}))
}

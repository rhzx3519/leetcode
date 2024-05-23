package main

import "fmt"

/**
https://leetcode.cn/problems/find-the-longest-equal-subarray/?envType=daily-question&envId=2024-05-23
*/
func longestEqualSubarray(nums []int, k int) int {
    m := make(map[int][]int)
    for i, x := range nums {
        m[x] = append(m[x], i)
    }
    var ans int
    for _, arr := range m {
        var l int
        for r := range arr {
            for arr[r]-arr[l]+1-(r-l+1) > k {
                l++
            }
            ans = max(ans, r-l+1)
        }
    }
    return ans
}

func main() {
    fmt.Println(longestEqualSubarray([]int{6, 4, 7, 6, 4, 8, 6, 6}, 1))
}

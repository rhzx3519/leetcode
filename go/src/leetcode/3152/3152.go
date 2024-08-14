package main

/**
https://leetcode.cn/problems/special-array-ii/?envType=daily-question&envId=2024-08-14
*/
func isArraySpecial(nums []int, queries [][]int) []bool {
    s := make([]int, len(nums))
    for i := 1; i < len(nums); i++ {
        s[i] = s[i-1]
        if nums[i-1]%2 == nums[i]%2 {
            s[i]++
        }
    }
    ans := make([]bool, len(queries))
    for i, q := range queries {
        ans[i] = s[q[0]] == s[q[1]]
    }
    return ans
}

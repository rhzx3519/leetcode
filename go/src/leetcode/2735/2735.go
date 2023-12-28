package main

/*
*
https://leetcode.cn/problems/collecting-chocolates/?envType=daily-question&envId=2023-12-28
*/
func minCost(nums []int, x int) int64 {
    sum := func(arr []int) int64 {
        var ans int64
        for _, i := range arr {
            ans += int64(i)
        }
        return ans
    }

    n := len(nums)
    f := make([]int, n)
    copy(f, nums)
    var ans = sum(nums)
    for k := 1; k < n; k++ {
        for i := range f {
            f[i] = min(f[i], nums[(i+k)%n])
        }
        ans = min(ans, int64(k)*int64(x)+sum(f))
    }

    return ans
}

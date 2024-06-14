package main

/**
https://leetcode.cn/problems/burst-balloons/?envType=daily-question&envId=2024-06-09
*/
func maxCoins(nums []int) int {
    n := len(nums)
    mem := make([][]int, n+1)
    for i := range mem {
        mem[i] = make([]int, n+1)
        for j := range mem[i] {
            mem[i][j] = -1
        }
    }
    var dfs func(int, int, int, int) int
    dfs = func(i int, j int, ll int, rr int) int {
        p := &mem[i][j]
        if *p != -1 {
            return *p
        }
        if i == j {
            *p = ll * nums[i] * rr
            return *p
        }
        var res int
        for k := i; k <= j; k++ {
            var left, right int
            if i < k {
                left = dfs(i, k-1, ll, nums[k])
            }
            if j > k {
                right = dfs(k+1, j, nums[k], rr)
            }
            res = max(res, left+right+ll*nums[k]*rr)
        }
        *p = res
        return *p
    }

    return dfs(0, n-1, 1, 1)
}

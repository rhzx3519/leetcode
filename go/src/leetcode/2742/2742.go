package main

import "math"

/**
https://leetcode.cn/problems/painting-the-walls/?envType=daily-question&envId=2024-06-28
@01背包 @动态规划
*/
func paintWalls(cost, time []int) int {
    n := len(cost)
    f := make([]int, n+1)
    for i := 1; i <= n; i++ {
        f[i] = math.MaxInt / 2 // 防止加法溢出
    }
    for i, c := range cost {
        t := time[i] + 1 // 注意这里加一了
        for j := n; j > 0; j-- {
            f[j] = min(f[j], f[max(j-t, 0)]+c)
        }
    }
    return f[n]
}

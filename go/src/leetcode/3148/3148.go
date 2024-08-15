package main

import "math"

/**
https://leetcode.cn/problems/maximum-difference-score-in-a-grid/?envType=daily-question&envId=2024-08-15
*/
func maxScore(grid [][]int) int {
    ans := math.MinInt
    m, n := len(grid), len(grid[0])
    f := make([][]int, m+1)
    f[0] = make([]int, n+1)
    for j := range f[0] {
        f[0][j] = math.MaxInt
    }
    for i, row := range grid {
        f[i+1] = make([]int, n+1)
        f[i+1][0] = math.MaxInt
        for j, x := range row {
            mn := min(f[i+1][j], f[i][j+1])
            ans = max(ans, x-mn)
            f[i+1][j+1] = min(mn, x)
        }
    }
    return ans
}

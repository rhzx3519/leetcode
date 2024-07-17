package main

import (
    "math"
    "math/bits"
)

/**
https://leetcode.cn/problems/number-of-possible-sets-of-closing-branches/?envType=daily-question&envId=2024-07-17
*/
func numberOfSets(n, maxDistance int, roads [][]int) int {
    g := make([][]int, n)
    for i := range g {
        g[i] = make([]int, n)
        for j := range g[i] {
            g[i][j] = math.MaxInt / 2 // 防止加法溢出
        }
    }
    for _, e := range roads {
        x, y, wt := e[0], e[1], e[2]
        g[x][y] = min(g[x][y], wt)
        g[y][x] = min(g[y][x], wt)
    }

    ans := 1 // s=0 一定满足要求
    f := make([][][]int, 1<<n)
    for i := range f {
        f[i] = make([][]int, n)
        for j := range f[i] {
            f[i][j] = make([]int, n)
            for k := range f[i][j] {
                f[i][j][k] = math.MaxInt / 2
            }
        }
    }
    f[0] = g
    for s := uint(1); s < 1<<n; s++ {
        k := bits.TrailingZeros(s)
        t := s ^ (1 << k)
        ok := true
        for i := 0; i < n; i++ {
            for j := 0; j < n; j++ {
                f[s][i][j] = min(f[t][i][j], f[t][i][k]+f[t][k][j])
                if ok && j < i && s>>i&1 != 0 && s>>j&1 != 0 && f[s][i][j] > maxDistance {
                    ok = false
                }
            }
        }
        if ok {
            ans++
        }
    }
    return ans
}

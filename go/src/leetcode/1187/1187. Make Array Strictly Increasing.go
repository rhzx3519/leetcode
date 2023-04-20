package main

import (
    "math"
    "sort"
)

/*
*
https://leetcode.cn/problems/make-array-strictly-increasing/
思路：动态规划 + 二分
f[i]表示不替换a[i]的情况下，最少操作次数
*/
func makeArrayIncreasing(a []int, b []int) int {
    sort.Ints(b)
    m := 0
    for _, x := range b[1:] { // b去重
        if b[m] != x {
            m++
            b[m] = x
        }
    }
    b = b[:m+1]

    a = append([]int{math.MinInt32}, append(a, math.MaxInt32)...)
    n := len(a)
    f := make([]int, n)
    for i := range f {
        f[i] = math.MaxInt32
    }
    f[0] = 0

    for i := 1; i < n; i++ {
        if a[i-1] < a[i] {
            f[i] = f[i-1]
        }
        j := sort.SearchInts(b, a[i])
        for k := 1; k <= min(i-1, j); k++ {
            if a[i-1-k] < b[j-k] {
                f[i] = min(f[i], f[i-1-k]+k)
            }
        }
    }

    if f[n-1] == math.MaxInt32 {
        return -1
    }
    return f[n-1]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

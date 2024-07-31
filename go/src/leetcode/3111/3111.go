package main

import "slices"

/**
https://leetcode.cn/problems/minimum-rectangles-to-cover-points/?envType=daily-question&envId=2024-07-31
*/
func minRectanglesToCoverPoints(points [][]int, w int) (ans int) {
    slices.SortFunc(points, func(p, q []int) int { return p[0] - q[0] })
    x2 := -1
    for _, p := range points {
        if p[0] > x2 {
            ans++
            x2 = p[0] + w
        }
    }
    return
}

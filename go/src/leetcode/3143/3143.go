package main

import (
    "math/bits"
    "sort"
)

func maxPointsInsideSquare(points [][]int, s string) (ans int) {
    sort.Search(1_000_000_001, func(size int) bool {
        vis := 0
        for i, p := range points {
            if abs(p[0]) <= size && abs(p[1]) <= size { // 点在正方形中
                c := s[i] - 'a'
                if vis>>c&1 > 0 { // c 在集合中
                    return true
                }
                vis |= 1 << c // 把 c 加入集合
            }
        }
        ans = bits.OnesCount(uint(vis))
        return false
    })
    return
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

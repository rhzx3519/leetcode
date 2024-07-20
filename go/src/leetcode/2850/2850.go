package main

import "math"

/**
https://leetcode.cn/problems/minimum-moves-to-spread-stones-over-grid/?envType=daily-question&envId=2024-07-20
*/
type pair struct{ x, y int }

func minimumMoves(grid [][]int) int {
    var from, to []pair
    for i, row := range grid {
        for j, cnt := range row {
            if cnt > 1 {
                for k := 1; k < cnt; k++ {
                    from = append(from, pair{i, j})
                }
            } else if cnt == 0 {
                to = append(to, pair{i, j})
            }
        }
    }

    ans := math.MaxInt
    permute(from, 0, func() {
        total := 0
        for i, f := range from {
            total += abs(f.x-to[i].x) + abs(f.y-to[i].y)
        }
        ans = min(ans, total)
    })
    return ans
}

func permute(a []pair, i int, do func()) {
    if i == len(a) {
        do()
        return
    }
    permute(a, i+1, do)
    for j := i + 1; j < len(a); j++ {
        a[i], a[j] = a[j], a[i]
        permute(a, i+1, do)
        a[i], a[j] = a[j], a[i]
    }
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

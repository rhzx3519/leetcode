package main

import "fmt"

/**
https://leetcode.cn/problems/maximum-number-of-moves-in-a-grid/?envType=daily-question&envId=2024-03-16
*/
var diff = [][]int{{-1, 1}, {0, 1}, {1, 1}}

func maxMoves(grid [][]int) (tot int) {
    m, n := len(grid), len(grid[0])
    f := make([][]int, m)
    for i := range f {
        f[i] = make([]int, n)
    }

    for j := n - 1; j >= 0; j-- {
        for i := 0; i < m; i++ {
            for _, d := range diff {
                ni, nj := i+d[0], j+d[1]
                if ni >= 0 && ni < m && nj >= 0 && nj < n && grid[ni][nj] > grid[i][j] {
                    f[i][j] = max(f[i][j], 1+f[ni][nj])
                }
            }
            if j == 0 {
                tot = max(tot, f[i][j])
            }
        }

    }
    return
}

func main() {
    fmt.Println(maxMoves([][]int{{2, 4, 3, 5}, {5, 4, 9, 3}, {3, 4, 2, 11}, {10, 9, 13, 15}}))
    fmt.Println(maxMoves([][]int{{3, 2, 4}, {2, 1, 9}, {1, 1, 7}}))
}

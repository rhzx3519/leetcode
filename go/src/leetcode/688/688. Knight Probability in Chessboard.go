/**
@author ZhengHao Lou
Date    2022/02/17
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/knight-probability-in-chessboard/
思路：记忆化搜索
*/
var next = [][]int{{-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {-1, -2}, {-2, -1}, {1, -2}, {2, -1}}

func knightProbability(n int, k int, row int, column int) float64 {
    var mem = make(map[[3]int]float64)
    var dfs func(r, c, k int) float64
    dfs = func(r, c, k int) float64 {
        var key = [3]int{r, c, k}
        if _, ok := mem[key]; ok {
            return mem[key]
        }

        if r < 0 || r >= n || c < 0 || c >= n {
            return 0
        }
        if k == 0 {
            return 1
        }
        var p float64
        for i := range next {
            nr, nc := r+next[i][0], c+next[i][1]
            p += dfs(nr, nc, k-1)
        }
        p /= 8.0
        mem[key] = p
        return p
    }
    return dfs(row, column, k)
}

func main() {
    fmt.Println(knightProbability(3, 2, 0, 0))
}

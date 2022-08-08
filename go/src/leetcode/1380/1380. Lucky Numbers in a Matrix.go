/**
@author ZhengHao Lou
Date    2022/02/15
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/lucky-numbers-in-a-matrix/
*/
func luckyNumbers(matrix [][]int) []int {
    m, n := len(matrix), len(matrix[0])
    rows, cols := make([]int, m), make([]int, n)
    for i := 0; i < m; i++ {
        var mi, c int = 1e6, -1
        for j := 0; j < n; j++ {
            if matrix[i][j] < mi {
                mi = matrix[i][j]
                c = j
            }
        }
        rows[i] = c
    }
    for j := 0; j < n; j++ {
        var mx, r int
        for i := 0; i < m; i++ {
            if matrix[i][j] > mx {
                mx = matrix[i][j]
                r = i
            }
            cols[j] = r
        }
    }

    var ans []int
    for i := range rows {
        c := rows[i]
        if cols[c] == i {
            ans = append(ans, matrix[i][c])
        }
    }
    fmt.Println(rows, cols)
    fmt.Println(ans)
    return ans
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    luckyNumbers([][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}})
}

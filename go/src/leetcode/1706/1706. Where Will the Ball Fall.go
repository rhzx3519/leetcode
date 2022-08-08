/**
@author ZhengHao Lou
Date    2022/02/24
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/where-will-the-ball-fall/
*/
func findBall(grid [][]int) []int {
    m, n := len(grid), len(grid[0])
    var ans = make([]int, n)
    for k := 0; k < n; k++ {
        var j = k
        for i := 0; i < m; i++ {
            if grid[i][j] == 1 {
                if j == n-1 || grid[i][j+1] == -1 {
                    j = -1
                    break
                } else {
                    j = j + 1
                }
            } else {
                if j == 0 || grid[i][j-1] == 1 {
                    j = -1
                    break
                } else {
                    j = j - 1
                }
            }
        }
        ans[k] = j
    }

    fmt.Println(ans)
    return ans
}

func main() {
    findBall([][]int{{1, 1, 1, -1, -1}, {1, 1, 1, -1, -1}, {-1, -1, -1, 1, 1}, {1, 1, 1, 1, -1}, {-1, -1, -1, -1, -1}})
    findBall([][]int{{-1}})
    findBall([][]int{{1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}, {1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}})
}

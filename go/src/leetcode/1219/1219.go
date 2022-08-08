/**
@author ZhengHao Lou
Date    2022/02/05
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/path-with-maximum-gold/
*/
var offset = [][]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}

func getMaximumGold(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    var mx int

    var dfs func(x, y, g int)
    dfs = func(x, y, g int) {
        if g > mx {
            mx = g
        }

        for _, dxy := range offset {
            nx, ny := x+dxy[0], y+dxy[1]
            if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] != 0 {
                t := grid[nx][ny]
                grid[nx][ny] = 0
                dfs(nx, ny, g+t)
                grid[nx][ny] = t
            }
        }
    }

    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] == 0 {
                continue
            }
            g := grid[i][j]
            grid[i][j] = 0
            dfs(i, j, g)
            grid[i][j] = g
        }
    }

    fmt.Println(mx)
    return mx
}

func main() {
    getMaximumGold([][]int{{0, 6, 0}, {5, 8, 7}, {0, 9, 0}})
}

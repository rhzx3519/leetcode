package main

import "fmt"

/*
*
https://leetcode.cn/problems/shortest-path-in-binary-matrix/
*/
type pair struct {
    x, y, step int
}

var diff = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

func shortestPathBinaryMatrix(grid [][]int) int {
    if grid[0][0] == 1 {
        return -1
    }
    n := len(grid)
    que := []pair{{0, 0, 1}}
    for len(que) != 0 {
        p := que[0]
        if p.x == n-1 && p.y == n-1 {
            return p.step
        }
        que = que[1:]
        for _, d := range diff {
            nx, ny := p.x+d[0], p.y+d[1]
            if nx >= 0 && nx < n && ny >= 0 && ny < n && grid[nx][ny] == 0 {
                grid[nx][ny] = 1
                que = append(que, pair{nx, ny, p.step + 1})
            }
        }
    }
    return -1
}

func main() {
    fmt.Println(shortestPathBinaryMatrix([][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}))
    fmt.Println(shortestPathBinaryMatrix([][]int{{0, 1}, {1, 0}}))
    fmt.Println(shortestPathBinaryMatrix([][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 0}}))
}

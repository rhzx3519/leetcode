package main

/**
https://leetcode.cn/problems/right-triangles/?envType=daily-question&envId=2024-08-02
*/
func numberOfRightTriangles(grid [][]int) (ans int64) {
    n := len(grid[0])
    colSum := make([]int, n)
    for _, row := range grid {
        for j, x := range row {
            colSum[j] += x
        }
    }

    for _, row := range grid {
        rowSum := -1 // 提前减一
        for _, x := range row {
            rowSum += x
        }
        for j, x := range row {
            if x == 1 {
                ans += int64(rowSum * (colSum[j] - 1))
            }
        }
    }
    return
}

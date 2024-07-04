package main

/**
https://leetcode.cn/problems/modify-the-matrix/?envType=daily-question&envId=2024-07-05
*/
func modifiedMatrix(matrix [][]int) [][]int {
    m, n := len(matrix), len(matrix[0])
    for j := 0; j < n; j++ {
        var mx int
        for i := 0; i < m; i++ {
            mx = max(mx, matrix[i][j])
        }
        for i := 0; i < m; i++ {
            if matrix[i][j] == -1 {
                matrix[i][j] = mx
            }
        }
    }
    return matrix
}

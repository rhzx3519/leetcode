/**
@author ZhengHao Lou
Date    2022/03/24
*/
package main

/**
https://leetcode-cn.com/problems/image-smoother/
*/
func imageSmoother(img [][]int) [][]int {
    m, n := len(img), len(img[0])
    smooth := make([][]int, m)
    for i := range smooth {
        smooth[i] = make([]int, n)
    }

    for i := range img {
        for j := range img[i] {
            var s, c int
            for k := max(0, i-1); k <= min(m-1, i+1); k++ {
                for z := max(0, j-1); z <= min(n-1, j+1); z++ {
                    s += img[k][z]
                    c++
                }
            }
            smooth[i][j] = s / c
        }
    }
    return smooth
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

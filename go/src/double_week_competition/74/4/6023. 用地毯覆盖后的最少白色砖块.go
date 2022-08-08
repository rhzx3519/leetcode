/**
@author ZhengHao Lou
Date    2022/03/19
*/
package main

/**
f[i][j][k]表示第i个砖块，剩余j条地毯时，上一个毯子位置为k时的最少的白色砖块数目
f[i][j][k] = f[i-1][j][k] if floor[i]=='0'

if floor[i]=='1'
    if k+carpetLen >= i
        f[i][j][k] = f[i-1][j][k]

*/
func minimumWhiteTiles(floor string, numCarpets int, carpetLen int) int {
    n := len(floor)
    pre := make([]int, n+1)
    for i := 0; i < n; i++ {
        var c int
        if floor[i] == '1' {
            c++
        }
        pre[i+1] = pre[i] + c
    }
    var total = pre[n]

    f := make([][][]int, n)
    for i := range f {
        f[i] = make([][]int, numCarpets)
        for j := range f[i] {
            f[i][j] = make([]int, n+1)
            for k := range f[i][j] {
                f[i][j][k] = pre[i+1]
            }
        }
    }
    for i := range floor {
        for j := numCarpets; j >= 0; j-- {
            if j == numCarpets {
                for k := 0; k < n; k++ {
                    f[i][j][n] = pre[i+1]
                }
            } else if j > 0 {
                for k := i - 1; k >= max(0, i-carpetLen); k-- {
                    f[i][j][n] = min(f[i][j][n], f[i][j][k])
                }
            } else {

            }

        }
    }

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

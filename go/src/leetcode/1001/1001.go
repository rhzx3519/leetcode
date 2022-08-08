/**
@author ZhengHao Lou
Date    2022/02/08
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/grid-illumination/
求解通过灯坐标的行直线与 xx 轴的交点，将交点的 xx 坐标作为通过灯坐标的行的数值。
求解通过灯坐标的列直线与 yy 轴的交点，将交点的 yy 坐标作为通过灯坐标的列的数值。
求解通过灯坐标的正对角线与 xx 轴的交点，将交点的 xx 坐标作为通过灯坐标的正对角线的数值。
求解通过灯坐标的反对角线与 yy 轴的交点，将交点的 yy 坐标作为通过灯坐标的反对角线的数值。
*/
func gridIllumination(n int, lamps [][]int, queries [][]int) []int {
    type pair struct {
        x, y int
    }
    points := map[pair]bool{}

    var rows = map[int]int{}
    var cols = map[int]int{}
    var diagonal = map[int]int{}
    var counterDiagonal = map[int]int{}
    for _, lamp := range lamps {
        x, y := lamp[0], lamp[1]
        p := pair{x: x, y: y}
        if points[p] {
            continue
        }
        points[p] = true
        rows[x]++
        cols[y]++
        diagonal[x-y]++
        counterDiagonal[x+y]++
    }

    ans := make([]int, len(queries))
    for i, q := range queries {
        x, y := q[0], q[1]
        if rows[x] > 0 || cols[y] > 0 || diagonal[x-y] > 0 || counterDiagonal[x+y] > 0 {
            ans[i] = 1
        }
        for nx := x - 1; nx <= x+1; nx++ {
            for ny := y - 1; ny <= y+1; ny++ {
                p := pair{x: nx, y: ny}
                if nx < 0 || nx >= n || ny < 0 || ny >= n || !points[p] {
                    continue
                }
                delete(points, p)
                rows[nx]--
                cols[ny]--
                diagonal[nx-ny]--
                counterDiagonal[nx+ny]--
            }
        }
    }
    fmt.Println(ans)
    return ans
}

func main() {
    gridIllumination(5, [][]int{{0, 0}, {4, 4}}, [][]int{{1, 1}, {1, 0}})
}

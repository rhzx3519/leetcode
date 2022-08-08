/**
@author ZhengHao Lou
Date    2022/04/24
*/
package main

import "fmt"

const N = 100001

func countLatticePoints(circles [][]int) int {
    var counter = map[int]int{}
    for _, c := range circles {
        x, y, r := c[0], c[1], c[2]
        for i := x - r; i <= x+r; i++ {
            for j := y - r; j <= y+r; j++ {
                if (i-x)*(i-x)+(j-y)*(j-y) <= r*r {
                    counter[encode(i, j)]++
                }
            }
        }
    }

    fmt.Println(len(counter))
    return len(counter)
}

func encode(i, j int) int {
    return i*N + j
}

func main() {
    //countLatticePoints([][]int{{2, 2, 1}})
    countLatticePoints([][]int{{2, 2, 2}, {3, 4, 1}})
}

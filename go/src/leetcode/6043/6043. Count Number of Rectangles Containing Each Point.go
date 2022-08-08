/**
@author ZhengHao Lou
Date    2022/04/24
*/
package main

import (
    "fmt"
    "sort"
)

const N = 100

func countRectangles(rectangles [][]int, points [][]int) []int {
    var counter = make([][]int, N+1)
    for _, rect := range rectangles {
        counter[rect[1]] = append(counter[rect[1]], rect[0])
    }
    for i := range counter {
        sort.Ints(counter[i])
    }

    var count = make([]int, len(points))

    for i, point := range points {
        x, y := point[0], point[1]
        var c int
        for j := y; j <= N; j++ {
            k := lowerBound(counter[j], x)
            c += len(counter[j]) - k
        }

        count[i] = c
    }
    fmt.Println(count)
    return count
}

func lowerBound(nums []int, x int) int {
    l, r := 0, len(nums)
    var m int
    for l < r {
        m = l + (r-l)>>1
        if nums[m] >= x {
            r = m
        } else {
            l = m + 1
        }
    }
    return l
}

func main() {
    countRectangles([][]int{{1, 2}, {2, 3}, {2, 5}}, [][]int{{2, 1}, {1, 4}})
    countRectangles([][]int{{1, 1}, {2, 2}, {3, 3}}, [][]int{{1, 3}, {1, 1}})
}

/**
@author ZhengHao Lou
Date    2022/04/07
*/
package main

import (
    "fmt"
    "sort"
)

/**
https://leetcode-cn.com/problems/find-players-with-zero-or-one-losses/
*/
func findWinners(matches [][]int) [][]int {
    var ans = make([][]int, 2)
    ind, out := map[int]int{}, map[int]int{}
    for _, match := range matches {
        ind[match[1]]++
        out[match[0]]++
    }
    for i := range out {
        if ind[i] == 0 {
            ans[0] = append(ans[0], i)
        }
    }
    for i := range ind {
        if ind[i] == 1 {
            ans[1] = append(ans[1], i)
        }
    }
    sort.Ints(ans[0])
    sort.Ints(ans[1])

    fmt.Println(ans)
    return ans
}

func main() {
    findWinners([][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}})
    findWinners([][]int{{2, 3}, {1, 3}, {5, 4}, {6, 4}})
}

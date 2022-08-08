/**
@author ZhengHao Lou
Date    2022/04/06
*/
package main

import (
    "fmt"
    "sort"
)

/**
https://leetcode-cn.com/problems/maximum-candies-allocated-to-k-children/
*/
func maximumCandies(candies []int, k int64) int {
    var s int64
    for i := range candies {
        s += int64(candies[i])
    }
    if s < k {
        return 0
    }

    sort.Ints(candies)
    n := int64(len(candies))

    l, r := 1, candies[n-1]+1
    var m int
    for l < r {
        m = l + (r-l)>>1
        var c int64
        for i := range candies {
            c += int64(candies[i] / m)
        }
        if c < k {
            r = m
        } else {
            l = m + 1
        }
    }

    fmt.Println(l, r, m)
    return l - 1
}

func main() {
    //maximumCandies([]int{5, 8, 6}, 5)
    //maximumCandies([]int{2, 5}, 3)
    maximumCandies([]int{1, 2, 3, 4, 10}, 5)
}

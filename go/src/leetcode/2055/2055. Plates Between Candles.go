/**
@author ZhengHao Lou
Date    2022/03/08
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/plates-between-candles/
*/
func platesBetweenCandles(s string, queries [][]int) []int {
    n := len(s)
    pre := make([]int, n+1)
    candles := []int{}
    for i, c := 1, 0; i <= n; i++ {
        switch s[i-1] {
        case '|':
            candles = append(candles, i-1)
            c++
        case '*':
        }
        pre[i] = c
    }
    fmt.Println(pre)
    fmt.Println(candles)
    var ans = make([]int, len(queries))
    for i := range queries {
        l, r := queries[i][0], queries[i][1]
        l1 := lowerBound(candles, l)
        r1 := upperBound(candles, r)
        if l1 == r1 {
            ans[i] = 0
            continue
        }
        // candles[l1] ~ candles[r1-1]
        c1, c2 := candles[l1], candles[r1-1]
        if c2 <= c1 {
            ans[i] = 0
            continue
        }
        between := pre[c2] - pre[c1+1]
        ans[i] = c2 - c1 - 1 - between
    }
    fmt.Println(ans)
    return ans
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

func upperBound(nums []int, x int) int {
    l, r := 0, len(nums)
    var m int
    for l < r {
        m = l + (r-l)>>1
        if nums[m] > x {
            r = m
        } else {
            l = m + 1
        }
    }
    return l
}

func main() {
    platesBetweenCandles("||*", [][]int{{2, 2}})
    //platesBetweenCandles("**|**|***|", [][]int{{2, 5}, {5, 9}})
    //platesBetweenCandles("***|**|*****|**||**|*", [][]int{{1, 17}, {4, 5}, {14, 17}, {5, 11}, {15, 16}})
}

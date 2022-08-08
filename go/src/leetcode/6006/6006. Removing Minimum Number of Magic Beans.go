/**
@author ZhengHao Lou
Date    2022/02/14
*/
package main

import (
    "fmt"
    "sort"
)

/**
https://leetcode-cn.com/problems/removing-minimum-number-of-magic-beans/
思路：贪心，
最优的策略是：升序排序
使所有非空袋子的豆子数目等于beans[i]的最少移除数目为, 将i之前的豆子全部移除 ，将i之后的豆子数目减少到beans[i]
time: O(nlogn)
space: O(1)
*/
func minimumRemoval(beans []int) int64 {
    n := len(beans)
    var total int64
    sort.Ints(beans)
    for _, b := range beans {
        total += int64(b)
    }

    var last int64
    var s = total
    for i := 0; i < n; i++ {
        total -= int64(beans[i])
        s = min(s, total-int64(beans[i])*int64(n-i-1)+last)
        last += int64(beans[i])
    }

    fmt.Println(s)
    return s
}

func min(a, b int64) int64 {
    if a < b {
        return a
    }
    return b
}

func main() {
    minimumRemoval([]int{4, 1, 6, 5})
    minimumRemoval([]int{2, 10, 3, 2})
}

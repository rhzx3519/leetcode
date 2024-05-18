package main

import (
    "fmt"
    "sort"
)

/**
https://leetcode.cn/problems/most-profit-assigning-work/submissions/532648061/?envType=daily-question&envId=2024-05-17
*/
func maxProfitAssignment(difficulty []int, profit []int, worker []int) (tot int) {
    n := len(difficulty)
    var index = make([]int, n)
    var maxProfit = make([]int, n)
    for i := 0; i < n; i++ {
        index[i] = i
    }
    sort.Slice(index, func(i, j int) bool {
        if difficulty[index[i]] != difficulty[index[j]] {
            return difficulty[index[i]] < difficulty[index[j]]
        }
        return profit[index[i]] < profit[index[j]]
    })
    var mx = 0
    for i := range index {
        mx = max(mx, profit[index[i]])
        maxProfit[i] = mx
    }
    for _, w := range worker {
        i := sort.Search(n, func(i int) bool {
            return difficulty[index[i]] >= w+1
        })
        if i > 0 {
            tot += maxProfit[i-1]
        }
    }
    fmt.Println(tot)
    return tot
}

func main() {
    maxProfitAssignment([]int{2, 4, 6, 8, 10}, []int{10, 20, 30, 40, 50}, []int{4, 5, 6, 7})
    maxProfitAssignment([]int{85, 47, 57}, []int{24, 66, 99}, []int{40, 25, 25})
    maxProfitAssignment([]int{68, 35, 52, 47, 86}, []int{67, 17, 1, 81, 3}, []int{92, 10, 85, 84, 82})
}

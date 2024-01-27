package main

import (
    "fmt"
    "slices"
    "sort"
)

/**
https://leetcode.cn/problems/maximum-number-of-alloys/?envType=daily-question&envId=2024-01-27
思路：二分
*/
func maxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) (tot int) {
    mx := slices.Min(stock) + budget
    for _, comp := range composition {
        tot += sort.Search(mx-tot, func(num int) bool {
            num += tot + 1
            var money int
            for i, s := range stock {
                if s < comp[i]*num {
                    money += (comp[i]*num - s) * cost[i]
                    if money > budget {
                        return true
                    }
                }
            }
            return false
        })
    }
    return
}

func main() {
    //fmt.Println(maxNumberOfAlloys(3, 2, 15, [][]int{{1, 1, 1}, {1, 1, 10}},
    //    []int{0, 0, 0}, []int{1, 2, 3}))
    fmt.Println(maxNumberOfAlloys(3, 2, 15, [][]int{{1, 1, 1}, {1, 1, 10}},
        []int{0, 0, 100}, []int{1, 2, 3}))
}

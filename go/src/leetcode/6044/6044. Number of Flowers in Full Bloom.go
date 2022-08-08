/**
@author ZhengHao Lou
Date    2022/04/24
*/
package main

import (
    "fmt"
    "sort"
)

func fullBloomFlowers(flowers [][]int, persons []int) []int {
    count := make([]int, len(persons))
    diff := make(map[int]int)
    for _, flower := range flowers {
        l, r := flower[0], flower[1]
        diff[l]++
        diff[r+1]--
    }
    var keys []int
    for k := range diff {
        keys = append(keys, k)
    }
    sort.Ints(keys)

    var nums = make(map[int]int)
    nums[keys[0]] = diff[keys[0]]
    for i := 1; i < len(keys); i++ {
        nums[keys[i]] = diff[keys[i]] + nums[keys[i-1]]
    }
    //fmt.Println(diff)
    fmt.Println(keys)
    for i := range keys {
        fmt.Print(nums[keys[i]])
        fmt.Print(", ")
    }
    fmt.Println()
    for i := range persons {
        j := upperBound(keys, persons[i])
        if j != 0 {
            count[i] = nums[keys[j-1]]
        }
    }

    fmt.Println(count)
    return count
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
    fullBloomFlowers([][]int{{1, 6}, {3, 7}, {9, 12}, {4, 13}}, []int{2, 3, 7, 11})
    fullBloomFlowers([][]int{{1, 10}, {3, 3}}, []int{3, 3, 2})
}

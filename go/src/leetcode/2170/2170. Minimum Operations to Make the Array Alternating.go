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
https://leetcode-cn.com/problems/minimum-operations-to-make-the-array-alternating/
*/
func minimumOperations(nums []int) int {
    var nEven, nOdd int
    var oddMap, evenMap = make(map[int]int), make(map[int]int)
    for i := range nums {
        if i&1 == 0 {
            nEven++
            evenMap[nums[i]]++
        } else {
            nOdd++
            oddMap[nums[i]]++
        }
    }

    var even, odd []int
    for x := range evenMap {
        even = append(even, x)
    }
    for x := range oddMap {
        odd = append(odd, x)
    }
    sort.Slice(even, func(i, j int) bool {
        return evenMap[even[i]] > evenMap[even[j]]
    })
    sort.Slice(odd, func(i, j int) bool {
        return oddMap[odd[i]] > oddMap[odd[j]]
    })

    n := len(nums)
    if n == 1 {
        return 0
    }
    if even[0] != odd[0] {
        return nEven - evenMap[even[0]] + nOdd - oddMap[odd[0]]
    }

    if nEven == evenMap[even[0]] {
        even = append(even, 0)
    }
    if nOdd == oddMap[odd[0]] {
        odd = append(odd, 0)
    }

    return min(nEven-evenMap[even[1]]+nOdd-oddMap[odd[0]],
        nEven-evenMap[even[0]]+nOdd-oddMap[odd[1]])

}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func main() {
    fmt.Println(minimumOperations([]int{3, 1, 3, 2, 4, 3}))
    //fmt.Println(minimumOperations([]int{1, 2, 2, 2, 2}))
}

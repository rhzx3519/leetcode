/**
@author ZhengHao Lou
Date    2022/04/11
*/
package main

import (
    "fmt"
    "sort"
)

/**
https://leetcode-cn.com/problems/largest-number-after-digit-swaps-by-parity/
*/
func largestInteger(num int) int {
    var even, odd []int
    var evenIdx, oddIdx []int
    var digits []int
    for i := 0; num != 0; i++ {
        digits = append(digits, num%10)
        num /= 10
    }
    reverse(digits)

    n := len(digits)
    for i := range digits {
        if digits[i]&1 == 0 {
            even = append(even, digits[i])
            evenIdx = append(evenIdx, i)
        } else {
            odd = append(odd, digits[i])
            oddIdx = append(oddIdx, i)
        }
    }
    sort.Ints(even)
    sort.Ints(odd)
    fmt.Println(even, odd)
    digits = make([]int, n)
    for i := range evenIdx {
        digits[evenIdx[i]] = even[len(even)-1]
        even = even[:len(even)-1]
    }

    for i := range oddIdx {
        digits[oddIdx[i]] = odd[len(odd)-1]
        odd = odd[:len(odd)-1]
    }
    fmt.Println(digits)
    var x int
    for i := range digits {
        x = x*10 + digits[i]
    }
    return x
}

func reverse(nums []int) {
    l, r := 0, len(nums)-1
    for l < r {
        nums[l], nums[r] = nums[r], nums[l]
        l++
        r--
    }
}

func main() {
    largestInteger(1234)
    //largestInteger(65875)
    //largestInteger(247) // 427
}

/**
@author ZhengHao Lou
Date    2022/04/07
*/
package main

import (
    "fmt"
    "strconv"
)

/**
https://leetcode-cn.com/problems/minimum-number-of-operations-to-convert-time/
*/
func convertTime(current string, correct string) int {
    s := toMiniutes(current)
    t := toMiniutes(correct)
    x := t - s
    var c int
    for _, inc := range []int{60, 15, 5, 1} {
        c += x / inc
        x %= inc
        if x == 0 {
            break
        }
    }

    fmt.Println(c)
    return c
}

func toMiniutes(t string) int {
    h, _ := strconv.Atoi(t[:2])
    m, _ := strconv.Atoi(t[3:])
    return h*60 + m
}

func main() {
    convertTime("02:30", "04:35")
    convertTime("11:00", "11:00")
}

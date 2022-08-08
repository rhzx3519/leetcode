/**
@author ZhengHao Lou
Date    2022/04/02
*/
package main

/**
https://leetcode-cn.com/contest/biweekly-contest-75/problems/minimum-bit-flips-to-convert-number/
*/
func minBitFlips(start int, goal int) int {
    x := start ^ goal
    var c int
    for x != 0 {
        c += x & 1
        x >>= 1
    }
    return c
}

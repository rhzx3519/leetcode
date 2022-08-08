/**
@author ZhengHao Lou
Date    2022/04/11
*/
package main

import (
    "fmt"
)

/**
https://leetcode-cn.com/problems/count-numbers-with-unique-digits/
*/
func countNumbersWithUniqueDigits(n int) int {
    var a, b = 1, 9
    for i := 1; i <= n; i++ {
        a += b
        b *= (10 - i)
    }
    return a
}

func main() {
    fmt.Println(countNumbersWithUniqueDigits(0))
}

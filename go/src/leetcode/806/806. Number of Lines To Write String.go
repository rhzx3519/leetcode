/**
@author ZhengHao Lou
Date    2022/04/12
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/number-of-lines-to-write-string/
*/
const CHUCK_SIZE = 100

func numberOfLines(widths []int, s string) []int {
    var total, line, c int
    for i := range s {
        x := widths[int(s[i]-'a')]
        total += x
        if c+x > CHUCK_SIZE {
            line++
            c = x
        } else {
            c += x
        }
    }
    if c > 0 {
        line++
    }
    return []int{line, c}
}

func main() {
    fmt.Println(numberOfLines([]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
        10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "abcdefghijklmnopqrstuvwxyz"))
}

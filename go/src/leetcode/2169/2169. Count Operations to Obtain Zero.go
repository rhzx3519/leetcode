/**
@author ZhengHao Lou
Date    2022/02/14
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/count-operations-to-obtain-zero/
*/
func countOperations(n1 int, n2 int) int {
    var c int
    for n1 != 0 && n2 != 0 {
        if n1 >= n2 {
            c += n1 / n2
            n1 %= n2
        } else {
            c += n2 / n1
            n2 %= n1
        }
    }
    fmt.Println(c)
    return c
}

func main() {
    countOperations(2, 3)
    countOperations(10, 10)
}

/**
@author ZhengHao Lou
Date    2022/02/25
*/
package main

import (
    "fmt"
    "strconv"
    "strings"
)

/**
https://leetcode-cn.com/problems/complex-number-multiplication/
*/
func complexNumberMultiply(num1 string, num2 string) string {
    x1, x2 := decompose(num1)
    y1, y2 := decompose(num2)
    z1 := x1*y1 - x2*y2
    z2 := x1*y2 + x2*y1
    return fmt.Sprintf("%v+%vi", z1, z2)
}

func decompose(num string) (int, int) {
    x := strings.Split(num, "+")
    x1, _ := strconv.Atoi(x[0])
    x2, _ := strconv.Atoi(x[1][:len(x[1])-1])
    return x1, x2
}

func main() {
    fmt.Println(complexNumberMultiply("1+1i", "1+1i"))
    fmt.Println(complexNumberMultiply("1+-1i", "1+-1i"))
    fmt.Println(complexNumberMultiply("1+-1i", "0+0i"))
}

/**
@author ZhengHao Lou
Date    2022/04/11
*/
package main

import (
    "fmt"
    "math"
    "strconv"
    "strings"
)

/**
https://leetcode-cn.com/problems/minimize-result-by-adding-parentheses-to-expression/
*/
func minimizeResult(expression string) string {
    var ans string
    var mi = math.MaxInt32
    plusIndex := strings.Index(expression, "+")
    for i := 0; i < plusIndex; i++ {
        for j := plusIndex + 2; j <= len(expression); j++ {
            var l, m, r int
            if i == 0 {
                l = 1
            } else {
                l, _ = strconv.Atoi(expression[:i])
            }
            if j == len(expression) {
                r = 1
            } else {
                r, _ = strconv.Atoi(expression[j:])
            }
            m1, _ := strconv.Atoi(expression[i:plusIndex])
            m2, _ := strconv.Atoi(expression[plusIndex:j])
            m = m1 + m2
            if mi > l*m*r {
                mi = l * m * r
                ans = fmt.Sprintf("%v(%v)%v", expression[:i], expression[i:j], expression[j:])
            }
        }
    }

    fmt.Println(ans)
    return ans
}

func main() {
    minimizeResult("247+38")
    minimizeResult("12+34")
    minimizeResult("999+999")
    minimizeResult("1+1") // "(1+1)"
}

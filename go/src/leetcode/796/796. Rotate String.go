/**
@author ZhengHao Lou
Date    2022/04/07
*/
package main

import "fmt"

func rotateString(s string, goal string) bool {
    if len(s) != len(goal) {
        return false
    }
    n := len(s)
    for i := 0; i < n; i++ {
        if s[:i] == goal[n-i:] && s[i:] == goal[:n-i] {
            return true
        }
    }
    return false
}

func main() {
    fmt.Println(rotateString("abcde", "cdeab"))
    fmt.Println(rotateString("abcde", "abced"))
}

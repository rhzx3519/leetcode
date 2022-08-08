/**
@author ZhengHao Lou
Date    2022/02/13
*/
package main

/**
https://leetcode-cn.com/problems/maximum-number-of-balloons/
*/
func maxNumberOfBalloons(text string) int {
    counter := make(map[byte]int)
    for i := range text {
        counter[text[i]]++
    }

    var c1, c2 int = 1e5, 1e5
    for _, b := range "ban" {
        c1 = min(c1, counter[byte(b)])
    }
    for _, b := range "lo" {
        c2 = min(c2, counter[byte(b)]>>1)
    }

    return min(c1, c2)
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

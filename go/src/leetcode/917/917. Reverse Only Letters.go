/**
@author ZhengHao Lou
Date    2022/02/23
*/
package main

import "unicode"

/**
https://leetcode-cn.com/problems/reverse-only-letters/
*/
func reverseOnlyLetters(s string) string {
    runes := []rune(s)
    for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
        for l < r && !unicode.IsLetter(rune(s[l])) {
            l++
        }
        for l < r && !unicode.IsLetter(rune(s[r])) {
            r--
        }
        if l < r {
            runes[l], runes[r] = runes[r], runes[l]
        }

    }
    return string(runes)
}

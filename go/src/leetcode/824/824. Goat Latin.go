/**
@author ZhengHao Lou
Date    2022/04/21
*/
package main

import (
    "strings"
    "unicode"
)

/**
https://leetcode-cn.com/problems/goat-latin/
*/
var vowels = map[rune]bool{
    'a': true,
    'e': true,
    'i': true,
    'o': true,
    'u': true,
}

func toGoatLatin(sentence string) string {
    words := strings.Split(sentence, " ")
    var goats []string
    for i, w := range words {
        var gw string
        if vowels[unicode.ToLower(rune(w[0]))] {
            gw = w + "ma" + strings.Repeat("a", i+1)
        } else {
            gw = w[1:] + string(w[0]) + "ma" + strings.Repeat("a", i+1)
        }
        goats = append(goats, gw)
    }
    return strings.Join(goats, " ")
}

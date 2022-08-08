/**
@author ZhengHao Lou
Date    2022/04/10
*/
package main

import (
    "fmt"
    "strings"
)

/**
https://leetcode-cn.com/problems/unique-morse-code-words/
*/
var (
    mose = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-",
        ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
)

func uniqueMorseRepresentations(words []string) int {
    counter := make(map[string]int)
    for _, w := range words {
        var tmp []string
        for i := range w {
            tmp = append(tmp, mose[int(w[i]-'a')])
        }
        counter[strings.Join(tmp, "")]++
    }
    fmt.Println(counter)
    return len(counter)
}

func main() {
    uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"})
}

/**
@author ZhengHao Lou
Date    2022/04/17
*/
package main

import (
    "fmt"
    "strings"
    "unicode"
)

/**
https://leetcode-cn.com/problems/most-common-word/
*/
func mostCommonWord(paragraph string, banned []string) string {
    var mx int
    var ans string
    ban := make(map[string]int)
    for _, b := range banned {
        ban[b]++
    }
    counter := make(map[string]int)
    n := len(paragraph)
    for i := 0; i < n; {
        var j = i
        for j < n && unicode.IsLetter(rune(paragraph[j])) {
            j++
        }
        if j > i {
            w := strings.ToLower(paragraph[i:j])
            if ban[w] == 0 {
                counter[w]++
                if counter[w] > mx {

                    mx = counter[w]
                    ans = w
                }
            }
            i = j
        } else {
            i++
        }
    }

    fmt.Println(counter)
    return ans
}

func main() {
    mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"})
}

package main

import (
    "fmt"
    "sort"
)

/*
*
https://leetcode.cn/problems/compare-strings-by-frequency-of-the-smallest-character/
*/
func numSmallerByFrequency(queries []string, words []string) []int {
    f := func(w string) int {
        var cnt = make([]int, 26)
        for i := range w {
            j := int(w[i] - 'a')
            cnt[j]++
        }

        var c int
        for j := range cnt {
            if cnt[j] != 0 {
                c = cnt[j]
                break
            }
        }
        return c
    }
    var ws []int
    for _, w := range words {
        ws = append(ws, f(w))
    }
    sort.Ints(ws)
    var ans []int
    for i := range queries {
        j := f(queries[i])
        k := sort.SearchInts(ws, j+1)
        ans = append(ans, len(ws)-k)
    }
    return ans
}

func main() {
    //fmt.Println(numSmallerByFrequency([]string{"cbd"}, []string{"zaaaz"}))
    //fmt.Println(numSmallerByFrequency([]string{"bbb", "cc"}, []string{"a", "aa", "aaa", "aaaa"}))
    fmt.Println(numSmallerByFrequency([]string{"bba", "abaaaaaa", "aaaaaa", "bbabbabaab", "aba", "aa", "baab", "bbbbbb", "aab", "bbabbaabb"},
        []string{"aaabbb", "aab", "babbab", "babbbb", "b", "bbbbbbbbab", "a", "bbbbbbbbbb", "baaabbaab", "aa"}))
}

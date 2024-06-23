package main

/**
https://leetcode.cn/problems/detect-capital/?envType=daily-question&envId=2024-06-23
*/
func detectCapitalUse(word string) bool {
    var b bool
    if word[0] >= 'A' && word[0] <= 'Z' {
        b = true
    }
    var b1 = 0
    var b2 = 0
    for _, c := range word[1:] {
        if c >= 'a' && c <= 'z' && b1 == 0 {
            b1 = 1
        }
        if c >= 'A' && c <= 'Z' && b2 == 0 {
            b2 = 1
        }
    }
    return (!b && b2 == 0) || (b && (b1 == 0 || b2 == 0))
}

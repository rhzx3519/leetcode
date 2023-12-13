package main

import "fmt"

/**
https://leetcode.cn/problems/lexicographically-smallest-palindrome/?envType=daily-question&envId=2023-12-13
*/
func makeSmallestPalindrome(s string) string {
    bytes := []byte(s)
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        if bytes[i] < bytes[j] {
            bytes[j] = bytes[i]
        } else if bytes[i] > bytes[j] {
            bytes[i] = bytes[j]
        }
    }
    return string(bytes)
}

func main() {
    fmt.Println(makeSmallestPalindrome("egcfe"))
}

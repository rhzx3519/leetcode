package main

/**
https://leetcode.cn/problems/lexicographically-smallest-string-after-substring-operation/?envType=daily-question&envId=2024-06-27
*/
func smallestString(s string) string {
    t := []byte(s)
    for i, c := range t {
        if c > 'a' {
            // 继续向后遍历
            for ; i < len(t) && t[i] > 'a'; i++ {
                t[i]--
            }
            return string(t)
        }
    }
    // 所有字母均为 a
    t[len(t)-1] = 'z'
    return string(t)
}

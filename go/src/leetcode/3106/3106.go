package main

/**
https://leetcode.cn/problems/lexicographically-smallest-string-after-operations-with-constraint/?envType=daily-question&envId=2024-07-27
*/
func getSmallestString(s string, k int) string {
    t := []byte(s)
    for i, c := range t {
        dis := int(min(c-'a', 'z'-c+1))
        if dis > k {
            t[i] -= byte(k)
            break
        }
        t[i] = 'a'
        k -= dis
    }
    return string(t)
}

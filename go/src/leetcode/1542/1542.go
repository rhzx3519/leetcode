package main

/**
https://leetcode.cn/problems/find-longest-awesome-substring/?envType=daily-question&envId=2024-05-20
*/
func longestAwesome(s string) int {
    prefix := map[int]int{0: -1}
    var seq int
    var ans int
    for i := range s {
        seq ^= 1 << int(s[i]-'0')
        if j, ok := prefix[seq]; ok {
            ans = max(ans, i-j)
        } else {
            prefix[seq] = i
        }

        for j := 0; j < 10; j++ {
            if k, ok := prefix[seq^(1<<j)]; ok {
                ans = max(ans, i-k)
            }
        }
    }
    return ans
}

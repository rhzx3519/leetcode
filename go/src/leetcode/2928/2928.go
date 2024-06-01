package main

/**
https://leetcode.cn/problems/distribute-candies-among-children-i/?envType=daily-question&envId=2024-06-01
*/
func distributeCandies(n int, limit int) int {
    var tot int
    for i := 0; i <= limit; i++ {
        for j := 0; j <= min(n-i, limit); j++ {
            if n-i-j <= limit {
                tot++
            }
        }
    }
    return tot
}

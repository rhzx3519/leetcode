package main

/**
https://leetcode.cn/problems/distribute-candies/?envType=daily-question&envId=2024-06-02
*/
func distributeCandies(candyType []int) int {
    n := len(candyType)
    genre := make(map[int]int)
    for i := range candyType {
        genre[candyType[i]]++
    }
    return min(len(genre), n/2)
}

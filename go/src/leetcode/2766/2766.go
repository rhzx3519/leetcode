package main

import "sort"

/**
https://leetcode.cn/problems/relocate-marbles/?envType=daily-question&envId=2024-07-24
*/
func relocateMarbles(nums []int, moveFrom []int, moveTo []int) []int {
    pos := make(map[int]int)
    for i := range nums {
        pos[nums[i]]++
    }
    for i := range moveFrom {
        if pos[moveFrom[i]] != 0 {
            pos[moveFrom[i]] = 0
            pos[moveTo[i]]++
        }
    }
    var ans []int
    for i := range pos {
        if pos[i] != 0 {
            ans = append(ans, i)
        }
    }
    sort.Ints(ans)
    return ans
}

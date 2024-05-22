package main

import "sort"

/**
https://leetcode.cn/problems/find-players-with-zero-or-one-losses/?envType=daily-question&envId=2024-05-22
*/
func findWinners(matches [][]int) [][]int {
    ind := map[int]int{}
    player := make(map[int]int)
    for i := range matches {
        a, b := matches[i][0], matches[i][1]
        player[a]++
        player[b]++
        ind[b]++
    }
    var answer = make([][]int, 2)
    for i := range player {
        if ind[i] == 0 {
            answer[0] = append(answer[0], i)
        } else if ind[i] == 1 {
            answer[1] = append(answer[1], i)
        }
    }
    sort.Ints(answer[0])
    sort.Ints(answer[1])
    return answer
}

package main

/**
https://leetcode.cn/problems/find-the-winner-of-an-array-game/?envType=daily-question&envId=2024-05-19
*/
func getWinner(arr []int, k int) int {
    var win = arr[0]
    var c int
    for i := 1; i < len(arr); i++ {
        if arr[i] > win {
            win = arr[i]
            c = 1
        } else {
            c++
        }
        if c >= k {
            break
        }
    }
    return win
}

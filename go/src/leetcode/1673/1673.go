package main

/**
https://leetcode.cn/problems/find-the-most-competitive-subsequence/?envType=daily-question&envId=2024-05-24
*/
func mostCompetitive(nums []int, k int) []int {
    nk := len(nums) - k
    var st []int
    for i := range nums {
        for nk != 0 && len(st) != 0 && st[len(st)-1] > nums[i] {
            nk--
            st = st[:len(st)-1]
        }
        st = append(st, nums[i])
    }

    return st[:k]
}

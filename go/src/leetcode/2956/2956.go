package main

/**
https://leetcode.cn/problems/find-common-elements-between-two-arrays/?envType=daily-question&envId=2024-07-16
*/
func findIntersectionValues(nums1 []int, nums2 []int) []int {
    var c1, c2 int
    m1, m2 := make(map[int]int), make(map[int]int)
    for i := range nums1 {
        m1[nums1[i]]++
    }
    for i := range nums2 {
        m2[nums2[i]]++
    }

    for i := range nums1 {
        if m2[nums1[i]] > 0 {
            c1++
        }
    }
    for i := range nums2 {
        if m1[nums2[i]] > 0 {
            c2++
        }
    }
    return []int{c1, c2}
}

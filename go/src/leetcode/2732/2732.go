package main

/**
https://leetcode.cn/problems/find-a-good-subset-of-the-matrix/?envType=daily-question&envId=2024-06-25
*/
func goodSubsetofBinaryMatrix(grid [][]int) []int {
    maskToIdx := map[int]int{}
    for i, row := range grid {
        mask := 0
        for j, x := range row {
            mask |= x << j
        }
        if mask == 0 {
            return []int{i}
        }
        maskToIdx[mask] = i
    }

    for x, i := range maskToIdx {
        for y, j := range maskToIdx {
            if x&y == 0 {
                return []int{min(i, j), max(i, j)}
            }
        }
    }
    return nil
}

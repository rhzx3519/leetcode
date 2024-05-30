package main

/**
https://leetcode.cn/problems/find-missing-and-repeated-values/?envType=daily-question&envId=2024-05-31
s1 = 2*a + ...
a+b+... = s
a - b = s1 - s
*/
func findMissingAndRepeatedValues(grid [][]int) []int {
    n := len(grid)

    cnt := make([]int, n*n+1)
    for i := range grid {
        for j := range grid[i] {
            cnt[grid[i][j]]++
        }
    }
    var a, b int
    for i := range cnt {
        if cnt[i] == 2 {
            a = i
        } else if cnt[i] == 0 {
            b = i
        }
    }
    return []int{a, b}
}

func main() {

}

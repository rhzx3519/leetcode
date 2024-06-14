package main

/**
https://leetcode.cn/problems/visit-array-positions-to-maximize-score/?envType=daily-question&envId=2024-06-14
*/
func maxScore(nums []int, x int) int64 {
    n := len(nums)
    f := make([][2]int, n+1)
    for i := n - 1; i >= 0; i-- {
        v := nums[i]
        r := v % 2
        f[i][r^1] = f[i+1][r^1] // v%2 != j 的情况
        f[i][r] = max(f[i+1][r], f[i+1][r^1]-x) + v
    }
    return int64(f[0][nums[0]%2])
}

func main() {
    maxScore([]int{2, 3, 6, 1, 9, 2}, 5)
}

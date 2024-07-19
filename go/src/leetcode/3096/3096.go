package main

/**
https://leetcode.cn/problems/minimum-levels-to-gain-more-points/?envType=daily-question&envId=2024-07-19
*/
func minimumLevels(possible []int) int {
    n := len(possible)
    // s = cnt1 - cnt0 = cnt1 - (n - cnt1) = cnt1 * 2 - n
    s := 0
    for _, x := range possible {
        s += x
    }
    s = s*2 - n
    pre := 0
    for i, x := range possible[:n-1] {
        pre += x*4 - 2
        if pre > s {
            return i + 1
        }
    }
    return -1
}

func main() {
    minimumLevels([]int{1, 0, 1, 0})
}

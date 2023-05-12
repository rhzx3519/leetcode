package main

import "math"

/*
*
https://leetcode.cn/problems/reverse-subarray-to-maximize-array-value/
作者：灵茶山艾府
链接：https://leetcode.cn/problems/reverse-subarray-to-maximize-array-value/solutions/2266500/bu-hui-hua-jian-qing-kan-zhe-pythonjavac-c2s6/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func maxValueAfterReverse(nums []int) int {
    base, d, n := 0, 0, len(nums)
    mx, mn := math.MinInt, math.MaxInt
    for i := 1; i < n; i++ {
        a, b := nums[i-1], nums[i]
        base += abs(a - b)
        mx = max(mx, min(a, b))
        mn = min(mn, max(a, b))
        d = max(d, max(abs(nums[0]-b)-abs(a-b), // i=0
            abs(nums[n-1]-a)-abs(a-b))) // j=n-1
    }
    return base + max(d, 2*(mx-mn))
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
func min(a, b int) int {
    if b < a {
        return b
    }
    return a
}
func max(a, b int) int {
    if b > a {
        return b
    }
    return a
}

/**
@author ZhengHao Lou
Date    2022/04/22
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/rotate-function/
a1[0]=nums[n-1], a1[1]=nums[0]..., f(1)=0*nums[n-1] + 1*nums[0] + 2*nums[1] + 3*nums[2] + .. + (n-1)*nums[n-2]
a0[0]=nums[0], a0[1]=nums[1]..., f(0)=0*nums[0] + 1*nums[1] + 2*nums[2] + .. + (n-1)*nums[n-1]
*/
func maxRotateFunction(nums []int) int {
    var mx int
    n := len(nums)
    var s, f0, f int
    for i := range nums {
        s += nums[i]
        f0 += i * nums[i]
    }
    mx = f0
    for i := 1; i < n; i++ {
        f = f0 + s - nums[n-i]*n
        mx = max(mx, f)
        f0 = f
    }

    fmt.Println(mx)
    return mx
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    maxRotateFunction([]int{4, 3, 2, 6})
    maxRotateFunction([]int{100})
}

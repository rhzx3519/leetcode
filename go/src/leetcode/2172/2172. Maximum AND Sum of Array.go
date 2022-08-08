/**
@author ZhengHao Lou
Date    2022/02/14
*/
package main

import "math/bits"

/**
https://leetcode-cn.com/problems/maximum-and-sum-of-array/
思路：状态压缩DP
可以用一个int表示所有篮子放了0、1、2个数字的状态
用2个bit位表示00, 放了0个数字；10/01放了1个数字，11放了两个数字
f[i]表示当篮子的状态为i时的最大和
*/
func maximumANDSum(nums []int, numSlots int) (ans int) {
    n := len(nums)
    f := make([]int, 1<<(2*numSlots))
    for i, fi := range f {
        c := bits.OnesCount(uint(i))
        if c >= n {
            continue
        }
        for j := 0; j < 2*numSlots; j++ {
            if i>>j&1 == 0 {
                s := i | 1<<j
                f[s] = max(f[s], fi+(j/2+1)&nums[c])
                ans = max(ans, f[s])
            }
        }
    }
    return
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

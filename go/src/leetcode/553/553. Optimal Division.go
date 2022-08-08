/**
@author ZhengHao Lou
Date    2022/02/27
*/
package main

import (
    "fmt"
    "strings"
)

/**
https://leetcode-cn.com/problems/optimal-division/
思路：除数最小的时候，结果最大
*/
func optimalDivision(nums []int) string {
    n := len(nums)
    if n == 1 {
        return fmt.Sprintf("%v", nums[0])
    }
    if n == 2 {
        return fmt.Sprintf("%v/%v", nums[0], nums[1])
    }
    var strs []string
    for _, num := range nums {
        strs = append(strs, fmt.Sprintf("%v", num))
    }

    return fmt.Sprintf("%v/(%v)", strs[0], strings.Join(strs[1:], "/"))
}

func main() {
    optimalDivision([]int{1000, 100, 10, 2})
}

/**
@author ZhengHao Lou
Date    2022/02/14
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/single-element-in-a-sorted-array/
*/
func singleNonDuplicate(nums []int) int {
    l, r := 0, len(nums)-1
    var m int
    for l < r {
        m = l + (r-l)>>1
        if nums[m] == nums[m-1] {
            if (m-l)&1 == 0 { // 0,1
                r = m - 2
            } else {
                l = m + 1
            }
        } else if nums[m] == nums[m+1] {
            if (r-m)&1 == 0 { // 0, 1
                l = m + 2
            } else {
                r = m - 1
            }
        } else {
            return nums[m]
        }
    }
    fmt.Println(nums[r])
    return nums[r]
}

func main() {
    singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8})
    singleNonDuplicate([]int{3, 3, 7, 7, 10, 11, 11})
    singleNonDuplicate([]int{1, 2, 2, 3, 3})
}

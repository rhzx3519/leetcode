/**
@author ZhengHao Lou
Date    2022/02/11
*/
package main

import (
    "math"
    "sort"
)

func minimumDifference(nums []int, k int) int {
    var ans = math.MaxInt32

    sort.Ints(nums)
    var low, high int
    for i := range nums {
        high = nums[i]
        if i >= k {
            low = nums[i-k+1]
        } else {
            low = nums[0]
        }

        if i+1 == k && high-low < ans {
            ans = high - low
        }

    }
    return ans
}

func main() {
    minimumDifference([]int{9, 4, 1, 7}, 2)
}

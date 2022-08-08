/**
@author ZhengHao Lou
Date    2022/03/05
*/
package main

import (
    "fmt"
    "sort"
)

func sortJumbled(mapping []int, nums []int) []int {
    var mp = func(x int) int {
        if x == 0 {
            return mapping[x]
        }
        var arr []int
        for ; x != 0; x /= 10 {
            arr = append(arr, mapping[x%10])
        }
        var y int
        for i := len(arr) - 1; i >= 0; i-- {
            y = y*10 + arr[i]
        }

        return y
    }
    mapTo := make(map[int]int)
    for _, num := range nums {
        mapTo[num] = mp(num)
    }

    sort.SliceStable(nums, func(i, j int) bool {
        return mapTo[nums[i]] < mapTo[nums[j]]
    })

    fmt.Println(nums)
    return nums
}

func main() {
    //sortJumbled([]int{8, 9, 4, 0, 2, 1, 3, 5, 7, 6}, []int{991, 338, 38})
    //sortJumbled([]int{8, 9, 4, 0, 2, 1, 3, 5, 7, 6}, []int{991, 338, 38})
    // [9,8,7,6,5,4,3,2,1,0]
    //sortJumbled([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
    // [454,95,404,47799,19021,162535,51890378]
    sortJumbled([]int{7, 9, 4, 1, 0, 3, 8, 6, 2, 5}, []int{47799, 19021, 162535, 454, 95, 51890378, 404})
}

/**
@author ZhengHao Lou
Date    2022/04/24
*/
package main

import "sort"

func intersection(nums [][]int) []int {
    counter := map[int]int{}
    for _, num := range nums[0] {
        counter[num]++
    }

    for i := 1; i < len(nums); i++ {
        var mp = make(map[int]int)
        for _, num := range nums[i] {
            if counter[num] > 0 {
                mp[num]++
            }
        }
        counter = mp
    }
    var ans []int
    for k := range counter {
        ans = append(ans, k)
    }
    sort.Ints(ans)
    return ans
}


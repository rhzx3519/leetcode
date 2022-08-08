/**
@author ZhengHao Lou
Date    2022/03/05
*/
package main

func mostFrequent(nums []int, key int) int {
    counter := make(map[int]int)
    for i := range nums {
        if nums[i] == key && i < len(nums)-1 {
            counter[nums[i+1]]++
        }
    }
    var target, c int
    for k, v := range counter {
        if v > c {
            target = k
            c = v
        }
    }
    return target
}

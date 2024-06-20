package main

/**
https://leetcode.cn/problems/6CE719/?envType=daily-question&envId=2024-06-21
*/
func temperatureTrend(temperatureA []int, temperatureB []int) int {
    n := len(temperatureA)
    trend := func(arr []int, i int) int {
        if arr[i] > arr[i+1] {
            return -1
        } else if arr[i] < arr[i+1] {
            return 1
        }
        return 0
    }
    var tot int
    var mx int
    for i := 0; i < n-1; i++ {
        if trend(temperatureA, i) == trend(temperatureB, i) {
            tot++
        } else {
            tot = 0
        }
        mx = max(mx, tot)
    }
    return mx
}

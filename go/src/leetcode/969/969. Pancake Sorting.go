/**
@author ZhengHao Lou
Date    2022/02/19
*/
package main

import "fmt"

func pancakeSort(arr []int) []int {
    var ans []int
    n := len(arr)
    for l := n; l > 0; l-- {
        var mx, idx int = 0, -1
        for i := 0; i < l; i++ {
            if arr[i] > mx {
                mx = arr[i]
                idx = i
            }
        }
        if idx == l-1 {
            continue
        }
        reverse(arr, 0, idx)
        ans = append(ans, idx+1)
        reverse(arr, 0, l-1)
        ans = append(ans, l)
    }

    fmt.Println(ans, arr)
    return ans
}

func reverse(arr []int, l, r int) {
    for l < r {
        arr[l], arr[r] = arr[r], arr[l]
        l++
        r--
    }
}

func main() {
    pancakeSort([]int{3, 2, 4, 1})
    pancakeSort([]int{1, 2, 3})
}

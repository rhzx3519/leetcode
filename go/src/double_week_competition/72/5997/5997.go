/**
@author ZhengHao Lou
Date    2022/02/19
*/
package main

func sumOfThree(num int64) []int64 {
    x := num / 3
    s := x - 1 + x + x + 1
    if s == num {
        return []int64{x - 1, x, x + 1}
    }
    if s+3 == num {
        return []int64{x, x + 1, x + 2}
    }
    if s-3 == num {
        return []int64{x - 2, x - 1, x}
    }
    return []int64{}
}

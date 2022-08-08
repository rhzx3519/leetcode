/**
@author ZhengHao Lou
Date    2022/04/09
*/
package main

/**
https://leetcode-cn.com/problems/reaching-points/
*/
func reachingPoints(sx int, sy int, tx int, ty int) bool {
    for tx != 0 && ty != 0 {
        if tx == sx && ty == sy {
            return true
        }
        if tx >= ty {
            tx -= max((tx-sx)/ty, 1) * ty
        } else {
            ty -= max((ty-sy)/tx, 1) * tx
        }
    }
    return false
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    reachingPoints(1, 1, 3, 5)
}

package main

import "fmt"

/*
*
https://leetcode.cn/problems/determine-the-winner-of-a-bowling-game/
*/
func isWinner(player1 []int, player2 []int) int {
    s1, s2 := score(player1), score(player2)
    fmt.Println(s1, s2)
    if s1 == s2 {
        return 0
    }
    if s1 > s2 {
        return 1
    }
    return 2
}

func score(a []int) int {
    var c, tot int
    for _, x := range a {
        if c > 0 {
            tot += x << 1
        } else {
            tot += x
        }

        if x == 10 {
            c = 2
        } else {
            c--
        }
    }
    return tot
}

func main() {
    isWinner([]int{0, 4, 7, 2, 0}, []int{2, 3, 3, 0, 1}) // 1
}

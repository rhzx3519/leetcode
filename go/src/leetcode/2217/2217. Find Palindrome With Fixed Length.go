/**
@author ZhengHao Lou
Date    2022/04/07
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/find-palindrome-with-fixed-length/
思路：只需要考虑构造前半部分，长度为n的回文数前半部分为10000 + k
*/
func kthPalindrome(queries []int, intLength int) []int64 {
    n := len(queries)
    ans := make([]int64, n)
    for i := range queries {
        ans[i] = kth(int64(queries[i]-1), intLength)
    }
    fmt.Println(ans)
    return ans
}

func kth(k int64, n int) int64 {
    if (n&1 == 0 && k >= 9*pow10(n>>1-1)) || (n&1 == 1 && k >= 9*pow10(n>>1)) {
        return -1
    }
    var x int64
    if n&1 == 0 {
        t := pow10(n>>1-1) + k
        x = t*pow10(n>>1) + reverse(t)
    } else {
        t := pow10(n>>1) + k
        x = t*pow10(n>>1) + reverse(t/10)
    }
    return x
}

func pow10(x int) int64 {
    var y = int64(1)
    for i := 0; i < x; i++ {
        y = y * 10
    }
    return y
}

func reverse(x int64) int64 {
    var y int64
    for ; x != 0; x /= 10 {
        y = y*10 + x%10
    }
    return y
}

func main() {
    //kthPalindrome([]int{1, 2, 3, 4, 5, 90}, 3)
    //kthPalindrome([]int{2, 4, 6}, 4)
    //[2,-1,8,-1,-1,-1,-1,9,7,6]
    kthPalindrome([]int{2, 201429812, 8, 520498110, 492711727, 339882032, 462074369, 9, 7, 6}, 1)
}

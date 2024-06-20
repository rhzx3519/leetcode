package main

/**
https://leetcode.cn/problems/number-of-beautiful-pairs/?envType=daily-question&envId=2024-06-20
*/
func countBeautifulPairs(nums []int) (tot int) {
    first := func(x int) int {
        var d = x
        for ; x != 0; x /= 10 {
            d = x % 10
        }
        return d
    }
    var gcd func(int, int) int
    gcd = func(a, b int) int {
        if b == 0 {
            return a
        }
        return gcd(b, a%b)
    }

    for i, a := range nums {
        for _, b := range nums[i+1:] {
            if gcd(first(a), b%10) == 1 {
                tot++
            }
        }
    }
    return
}

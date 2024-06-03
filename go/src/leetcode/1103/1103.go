package main

func distributeCandies(candies int, n int) []int {
    var ans = make([]int, n)
    for i := 1; candies > 0; i++ {
        ans[(i-1)%n] += min(i, candies)
        candies -= i
    }
    return ans
}

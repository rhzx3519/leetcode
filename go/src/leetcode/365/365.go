package main

/**
https://leetcode.cn/problems/water-and-jug-problem/?envType=daily-question&envId=2024-01-28
*/
func canMeasureWater(x int, y int, z int) bool {
    if z == 0 {
        return true
    }
    if x == 0 {
        return y == z
    }
    if y == 0 {
        return x == z
    }
    if x+y < z {
        return false
    }
    return z%gcd(x, y) == 0
}

func gcd(a, b int) int {
    if b == 0 {
        return a
    }
    return gcd(b, a%b)
}

/**
思路1：转化为01背包问题，求一个数字能够由[0, 3^1, 3^2, ..]这样的数组加和组成

思路2：如果一个数 可以表示成若干个不同的三的幂之和，它的三进制数只能由0和1组成
 */
package main

import (
	"fmt"
	"math"
)

func checkPowersOfThree(n int) bool {
	nums := []int{}
	for i := 0; i <= 14; i++ {
		nums = append(nums, int(math.Pow(3, float64(i))))
	}

	dp := make([]bool, n+1)
	dp[0] = true

	for _, num := range nums {
		for i := n; i >= num; i-- {
			dp[i] = dp[i] || dp[i - num]
		}
	}

	return dp[n]
}

func checkPowersOfThree1(n int) bool {
	for n != 0 {
		if n%3 == 2 {
			return false
		}
		n /= 3
	}
	return true
}

func main() {
	fmt.Println(checkPowersOfThree(12))
	fmt.Println(checkPowersOfThree(91))
	fmt.Println(checkPowersOfThree(21))
	fmt.Println(checkPowersOfThree(7814780))
}

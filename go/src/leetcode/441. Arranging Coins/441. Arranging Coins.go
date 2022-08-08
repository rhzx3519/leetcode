/**
@author ZhengHao Lou
Date    2021/10/10
*/
package main

import "math"

func arrangeCoins(n int) int {
	return int((math.Sqrt(float64(1 + 8*n)) - 1) * 0.5)
}

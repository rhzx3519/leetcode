/**
@author ZhengHao Lou
Date    2021/11/30
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/nth-digit/
c1: 1*10 - 1	1~9 = 1*9 = 9
c2: 2*100 - c1	10~99 = 20*9 = 180
c3: 3*1000 - c2	 100~999 = 300*9 = 2700
 */
func findNthDigit(n int) int {
	var i, d, c = 1, 1, 1
	var x = 1
	for ; ; x = x*10 {
		c = d*x*9
		if i + c >= n {
			break
		}
		i += c
		d++
	}

	fmt.Println(i, c, x)
	j := (n - i) / d
	t := (n - i) % d
	y := x + j
	fmt.Println(j, t, y)
	var digits = []int{}
	for y != 0 {
		digits = append(digits, y % 10)
		y /= 10
	}
	fmt.Println(digits[len(digits) - 1 - t])
	return digits[len(digits) - 1 - t]

}

func main() {
	//findNthDigit(3)
	findNthDigit(11)
	findNthDigit(1000)	// 3
}

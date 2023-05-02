/*
*
@author ZhengHao Lou
Date    2021/12/02
*/
package main

import (
	"fmt"
	"math"
)

/*
*
https://leetcode-cn.com/problems/powerful-integers/
*/
func powerfulIntegers(x int, y int, bound int) []int {
	counter := make(map[int]int)
	for i := 0; pow(x, i) <= bound-1; i++ {
		for j := 0; pow(y, j) <= bound-pow(x, i); j++ {
			t := pow(x, i) + pow(y, j)
			counter[t]++
			if y == 1 {
				break
			}
		}
		if x == 1 {
			break
		}
	}

	ans := []int{}
	for i, _ := range counter {
		ans = append(ans, i)
	}
	return ans
}

func pow(x, i int) int {
	return int(math.Pow(float64(x), float64(i)))
}

func main() {
	//fmt.Println(powerfulIntegers(2, 3, 10))
	//fmt.Println(powerfulIntegers(3, 5, 15))
	fmt.Println(powerfulIntegers(2, 1, 10))
}

/**
@author ZhengHao Lou
Date    2021/11/17
*/
package main

import (
	"fmt"
	"math"
)

/**
https://leetcode-cn.com/problems/koko-eating-bananas/
思路：二分
K的边界为[1, max(piles)]
time: O(log(max(piles)) n)
space: O(1)
 */
func minEatingSpeed(piles []int, h int) int {
	var l, r = 1, maxEle(piles)
	var k, c int
	for l < r {
		k = l + (r - l)>>1
		c = 0
		for _, p := range piles {
			c += int(math.Ceil(float64(p)/ float64(k)))
		}

		//fmt.Println(l, r, k, c)
		if c <= h {
			r = k
		} else {
			l = k + 1
		}
	}
	//fmt.Println(l, r, c)
	return l
}

func maxEle(piles []int) int {
	var y int
	for _, x := range piles {
		if x > y {
			y = x
		}
	}
	return y
}

func main() {
	fmt.Println(minEatingSpeed([]int{3,6,7,11}, 8))
	fmt.Println(minEatingSpeed([]int{30,11,23,4,20}, 5))
	fmt.Println(minEatingSpeed([]int{30,11,23,4,20}, 6))
}

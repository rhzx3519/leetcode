/**
@author ZhengHao Lou
Date    2021/12/17
*/
package main

/**
https://leetcode-cn.com/problems/water-bottles/
 */
func numWaterBottles(numBottles int, numExchange int) int {
	var c = numBottles
	var i, r, t int
	for i = numBottles; i + r >= numExchange; {
		t = i + r
		i = t / numExchange	// bottles of wine in next round
		r = t % numExchange	// empty bottles
		c += i
	}
	//fmt.Println(c)
	return c
}

func main() {
	numWaterBottles(9, 3)
	numWaterBottles(15, 4)
	numWaterBottles(5, 5)
	numWaterBottles(2, 3)
	numWaterBottles(15, 8)
}

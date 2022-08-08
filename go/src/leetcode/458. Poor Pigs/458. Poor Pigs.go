/**
@author ZhengHao Lou
Date    2021/11/25
*/
package main

import (
	"fmt"
	"math"
)

/**
https://leetcode-cn.com/problems/poor-pigs/
思路：数学题
一只猪携带的信息量为times = minutesToTest / minutesToDie + 1
假设minutesToTest = 60, minutesToDie = 15, times = 60/15+1 = 5，具体的，
1只猪,
0~15 喝0号桶
15~30 喝1号桶
30~45 喝2号桶
45~60 喝3号桶
不喝4号桶，
可以验证5只桶.

如果是两只猪，则可以验证25只桶，具体的，按行列分成5x5的矩阵，pig1按行喝混合水，pig2按列喝混合水，
如果某个时间段猪死了，则按照行列可以确定唯一一桶有毒的水
依此类推，3只猪可以验证5^3只桶
 */
func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	times := int(math.Ceil(float64(minutesToTest) / float64(minutesToDie)) + 1)
	var x, y = 0, 1
	for y < buckets {
		y *= times
		x++
	}
	return x
}

func main() {
	fmt.Println(poorPigs(1000, 15, 60))
	fmt.Println(poorPigs(5, 15, 60))
	fmt.Println(poorPigs(4, 15, 15))
	fmt.Println(poorPigs(4, 15, 30))
}

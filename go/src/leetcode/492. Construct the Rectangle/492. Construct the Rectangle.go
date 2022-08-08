/**
@author ZhengHao Lou
Date    2021/10/23
*/
package main

import "math"

func constructRectangle(area int) []int {
	x := int(math.Sqrt(float64(area)))
	for area %x != 0 && x > 0{
		x--
	}
	return []int{area/x, x}
}

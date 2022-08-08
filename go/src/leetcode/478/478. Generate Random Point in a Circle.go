/**
@author ZhengHao Lou
Date    2022/06/05
*/
package main

import (
	"math"
	"math/rand"
)

type Solution struct {
	r, x, y float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{
		r: radius,
		x: x_center,
		y: y_center,
	}
}

func (this *Solution) RandPoint() []float64 {
	theta := math.Pi * 2 * rand.Float64()
	r := math.Sqrt(rand.Float64() * this.r * this.r)
	return []float64{math.Sin(theta)*r + this.x, math.Cos(theta)*r + this.y}
}

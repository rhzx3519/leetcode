package main

import (
	"fmt"
	"math"
)

func countTriples(n int) int {
	var count int
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			c2 := i*i+j*j
			c := math.Sqrt(float64(c2))
			t := int(c);
			if t*t == c2 && t <= n{
				//fmt.Println(i, j, t)
				count++
			}

		}
	}
	return count
}

func main() {
	fmt.Println(countTriples(5))
	fmt.Println(countTriples(10))
}
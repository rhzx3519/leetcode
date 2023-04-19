package main

import "fmt"

/*
*
返回x的所有质因数, x>=2
*/
func primeFactor(x int) []int {
	var factors []int
	t := x
	for d := 2; d*d <= t; d++ {
		if t%d == 0 {
			factors = append(factors, d)
			for t /= d; t%d == 0; t /= d {
			}
		}
	}

	if t > 1 {
		factors = append(factors, t)
	}
	return factors
}

func main() {
	fmt.Println(primeFactor(16))
}

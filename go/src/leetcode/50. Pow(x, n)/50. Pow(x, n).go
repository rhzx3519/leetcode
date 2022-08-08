package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return float64(1)
	}

	if n == 1 {
		return x
	}

	if n < 0 {
		return 1.0 / myPow(x, -n)
	}

	half := myPow(x, n>>1)
	if n & 1 == 1 {
		return half * half * x
	}
	return half * half
}

func main() {
	fmt.Println(myPow(2.0, 10))
	fmt.Println(myPow(2.1, 3))
	fmt.Println(myPow(2.0, -2))
}

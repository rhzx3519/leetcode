package main

import "fmt"

func clumsy(N int) int {
	ans := 0
	sign := 1
	for ; N /4 != 0; N -= 4 {
		ans += sign*N*(N-1)/(N-2) + (N-3)
		sign = -1
	}
	switch N {
	case 3:
		ans += sign*N*(N-1)/(N-2)
	case 2:
		ans += sign*N*(N-1)
	case 1:
		ans += sign*N
	}
	return ans
}

func main() {
	fmt.Println(clumsy(11))
}
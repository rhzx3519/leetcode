package main

import "fmt"

func climbStairs(n int) int {
	var f1, f2 = 1, 1
	var f3 = 0
	for i := 2; i <= n; i++ {
		f3 = f1 + f2
		f1 = f2
		f2 = f3
	}
	return f2
}

func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
}

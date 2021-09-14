package main

import (
	"fmt"
	"math"
)

var mod = int(math.Pow10(9)) + 7
func fib(n int) int {
	c := make(chan int)
	quit := make(chan int)
	var ans int
	go func() {
		for i := 0; i <= n; i++ {
			//fmt.Println(<- c)
			ans = (<- c) % mod
		}
		quit <- 0
	}()
	
	doFib(c, quit)
	return ans
}

func doFib(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, (x+y)%mod
		case <-quit:
			return
		}
	}
}

func main() {
	fmt.Println(fib(0))
	fmt.Println(fib(1))
	fmt.Println(fib(2))
	fmt.Println(fib(5))
	fmt.Println(fib(45))
	fmt.Println(fib(95))
}

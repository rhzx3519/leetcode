package main

import "fmt"

func minimumPerimeter(neededApples int64) int64 {
	var apples int64
	var x = int64(2)
	for apples < neededApples {
		apples += (x + x*2)*2
		apples += sum(x)*4
		apples += sum(x-1)*4
		apples += (2*x-1)*2*x
		x++
	}
	fmt.Println(x*4)
	return x*4
}

func sum(x int64) int64 {
	return (1 + x)*x/2
}

func main() {
	//minimumPerimeter(1)
	minimumPerimeter(13)
	//minimumPerimeter(1000000000)
}
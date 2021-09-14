package main

import "fmt"

/**
sn = 2n(n+1)(2n+1)
 */
func minimumPerimeter(neededApples int64) int64 {
	var n = int64(1)
	for 2*n*(n+1)*(2*n+1) < neededApples {
		n++
	}
	return n*8
}


func main() {
	fmt.Println(minimumPerimeter(1))
}

package main

import (
	"fmt"
)

func reverse(x int) int {
	t := int64(0)
	for x != 0 {
		t = t*10 + int64(x%10)
		x /= 10
	}
	if int64(int32(t)) == t {
		return int(t)
	}
	return 0
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
	fmt.Println(reverse(0))
}

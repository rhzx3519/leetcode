package main

import "fmt"

func reinitializePermutation(n int) int {
	step, pos := 0, 1
	for {
		if pos&1 == 0 {
			pos /= 2
		} else {
			pos = n/2 + (pos-1)/2
		}
		step++

		if pos == 1 {
			break
		}
	}

	return step
}

func main() {
	fmt.Println(reinitializePermutation(2))
	fmt.Println(reinitializePermutation(4))
	fmt.Println(reinitializePermutation(6))
}
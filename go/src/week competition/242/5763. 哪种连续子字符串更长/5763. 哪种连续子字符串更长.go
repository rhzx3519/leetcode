package main

import "fmt"

func checkZeroOnes(s string) bool {
	var m1, m0 int
	var c1, c0 int
	for _, c := range s {
		if c == '1' {
			c0 = 0
			c1++
		} else {
			c1 = 0
			c0++
		}
		m1 = max(c1, m1)
		m0 = max(c0, m0)
	}

	return m1 > m0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(checkZeroOnes("01111110"))
}
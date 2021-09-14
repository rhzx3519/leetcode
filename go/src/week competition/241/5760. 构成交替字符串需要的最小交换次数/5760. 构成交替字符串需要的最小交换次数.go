package main

import "fmt"

func minSwaps(s string) int {
	n := len(s)

	var try = func(f byte) int {
		c := [2]int{}

		for i := 0; i < n; i++ {
			if s[i] != f {
				if f == '1' {
					c[1]++
				} else {
					c[0]++
				}
			}


			if f == '1' {
				f = '0'
			} else {
				f = '1'
			}
		}
		if c[0] != c[1] {
			return -1
		}
		return c[0]
	}

	r1 := try('0')
	r2 := try('1')
	if r1 == -1 && r2 == -1 {
		return -1
	} else if r1 == -1 {
		return r2
	} else if r2 == -1 {
		return r1
	}
	return min(r1, r2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minSwaps("111000"))
	fmt.Println(minSwaps("010"))
	fmt.Println(minSwaps("1110"))
	fmt.Println(minSwaps("100"))
}


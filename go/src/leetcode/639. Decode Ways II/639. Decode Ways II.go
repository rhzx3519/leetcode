package main

import "fmt"

/**
https://leetcode-cn.com/problems/decode-ways-ii/
time: O(n)
space: O(n)
 */
const mod int = 1e9 +7

func numDecodings(s string) int {
	n := len(s)
	f := make([]int, n + 1)
	f[0] = 1

	for i := 1; i <= n; i++ {
		c := s[i - 1]
		if c == '*' {
			f[i] = add(f[i], f[i-1]*9)
		} else if c >= '1' && c <= '9' {
			f[i] = add(f[i], f[i-1])
		}
		if i > 1 {
			c1 := s[i-2]
			if c1 == '1' {
				if c == '*' {
					f[i] = add(f[i], f[i-2]*9)
				} else {
					f[i] = add(f[i], f[i-2])
				}
			} else if c1 == '2' {
				if c == '*' {
					f[i] = add(f[i], f[i-2]*6)
				} else if c >= '0' && c <= '6' {
					f[i] = add(f[i], f[i-2])
				}
			} else if c1 == '*' {
				if c == '*' {
					f[i] = add(f[i], f[i-2]*15)
				} else if c >= '0' && c <= '6'{		// "*0", 10, 20
					f[i] = add(f[i], f[i-2]*2)
				} else {							// "*7", 17
					f[i] = add(f[i], f[i-2])
				}
			}
		}
	}

	// fmt.Println(f)
	return f[n] % mod
}

func add(a, b int) int {
	return (a + b) % mod
}
func main() {
	fmt.Println(numDecodings("*"))
	fmt.Println(numDecodings("1*"))
	fmt.Println(numDecodings("2*"))		// 15
	fmt.Println(numDecodings("**"))		// 96
	fmt.Println(numDecodings("*1"))		// 11
	fmt.Println(numDecodings("*1*1*0"))	// 404
	fmt.Println(numDecodings("*********"))	// 291868912
}

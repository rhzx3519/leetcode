package main

import (
	"fmt"
)

func maxProduct(s string) int {
	n := len(s)
	valid := make([]bool, 1<<n)
	f := make([]int, 1<<n)	// f[mask]表示当状态为mask时，最大的回文序列长度
	f[0] = 0

	//var mask int
	for mask := 1; mask < 1<<n; mask++ {
		var is = true
		var sz int
		l, r := 0, n-1
		for l <= r {
			for l <= r && mask & (1<<l) == 0 {
				l++
			}
			for l <= r && mask & (1<<r) == 0 {
				r--
			}

			if l <= r {
				if s[l] == s[r] {
					if l != r {
						sz += 2
					} else {
						sz++
					}
				} else {
					is = false
					break
				}
			}

			l++
			r--
		}

		if sz > 0 && is == true {
			valid[mask] = true
			f[mask] = sz

		}
	}
	for mask := 1; mask < 1<<n; mask++ {
		if valid[mask] {
			//fmt.Printf("%011b\n", mask)
		}
	}


	var ans int
	for mask := 1; mask < 1<<n; mask++ {
		if !valid[mask] {
			continue
		}
		fmt.Printf("%011b\n", mask)
		t := (1<<n - 1) ^ mask 	// e.g. 010 -> 111^010 -> 101
		for subset := t; subset != 0; subset = (subset - 1)&t {
			if !valid[subset] {
				continue
			}
			fmt.Printf("%011b\n", subset)
			ans = max(ans, f[subset] * f[mask])
		}

	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//fmt.Println(maxProduct("bb"))
	//fmt.Println(maxProduct("leetcodecom"))
	//fmt.Println(maxProduct("accbcaxxcxx"))
	fmt.Println(maxProduct("zqz"))
}


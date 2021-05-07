package main

import (
	"fmt"
	"strconv"
)

func numDecodings(s string) int {
	var f1, f2, f3 = 0, 1, 0
	for i, c := range s {
		f3 = 0
		if (c - '0') > 0 && (c - '0') <= 26 {
			f3 = f2
		}
		if i > 0 && s[i-1] != '0' {
			if n, _ := strconv.Atoi(s[i-1:i+1]); n <= 26 && n >= 1 {
				f3 += f1
			}
		}

		f1 = f2
		f2 = f3
	}

	return f3
}

func main() {
	fmt.Println(numDecodings("10"))
	//fmt.Println(numDecodings("12"))
	//fmt.Println(numDecodings("225"))
	//fmt.Println(numDecodings("0"))
	//fmt.Println(numDecodings("06"))
}

package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	n := len(s)
	i := 0
	for ; i < n && s[i]==' '; i++ {}
	if i == n {
		return 0
	}

	sign := true
	if s[i] == '+' {
		i++;
	} else if s[i] == '-' {
		sign = false
		i++;
	}
	t := uint64(0)
	for ; i < n && s[i] >= '0' && s[i] <= '9'; i++ {
		t = t*10 + uint64(s[i] - '0')
		if sign && t > math.MaxInt32 {
			return math.MaxInt32
		} else if !sign && t > 1<<31 {
			return math.MinInt32
		}
	}




	if !sign {
		return -int(t)
	}

	return int(t)
}

func main() {
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi("   -42"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("words and 987"))
	fmt.Println(myAtoi("-91283472332"))
	fmt.Println(myAtoi("9223372036854775808"))
	fmt.Println(myAtoi("-2147483649"))
	fmt.Println(myAtoi("18446744073709551617"))
}

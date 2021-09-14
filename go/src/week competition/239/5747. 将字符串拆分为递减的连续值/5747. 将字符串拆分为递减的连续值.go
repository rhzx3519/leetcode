package main

import (
	"fmt"
	"math"
	"strconv"
)

func splitString(s string) bool {
	return dp(0, 0, math.MaxInt32, s)
}

func dp(idx int, cnt int, last int, s string) bool {
	if idx == len(s) {
		return cnt > 1
	}
	for i := idx; i < len(s); i++ {
		num, _ := strconv.Atoi(s[idx:i+1])
		if last != math.MaxInt32 && num - last != -1 {
			continue
		}
		if dp(i+1, cnt+1, num, s) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(splitString("12"))
	fmt.Println(splitString("050043"))
	fmt.Println(splitString("9080701"))
	fmt.Println(splitString("10009998"))
}
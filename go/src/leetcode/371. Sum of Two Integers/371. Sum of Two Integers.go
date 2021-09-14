package main

import "fmt"

func getSum(a int, b int) int {
	var sum, carry int
	sum = a ^ b			// 不进位之和
	carry = (a & b) << 1	// 进位
	if carry != 0 {
		return getSum(sum, carry)
	}
	return sum
}

func main() {
	fmt.Println(getSum(1, 2))
	fmt.Println(getSum(-2, 3))
}

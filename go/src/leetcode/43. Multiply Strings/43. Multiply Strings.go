package main

import "fmt"

func multiply(num1 string, num2 string) string {
	n1, n2 := len(num1), len(num2)
	c := make([]int, n1 + n2)
	for i := n1 - 1; i >= 0; i-- {
		for j := n2 - 1; j >= 0; j-- {
			c[i + j + 1] += int(num1[i] - '0') * int(num2[j] - '0')	// 从低位往高位相乘
		}
	}
	var i = n1 + n2 - 1
	for i > 0 {		// 从低位往高位进位
		c[i-1] += c[i] / 10
		c[i] %= 10
		i--
	}

	i = 0
	for i < n1 + n2 && c[i] == 0 {
		i++
	}
	if i == n1 + n2 {
		return "0"
	}

	bytes := []byte{}
	for ; i < n1 + n2; i++ {
		bytes = append(bytes, byte(c[i] + '0'))
	}
	return string(bytes)
}

func main() {
	fmt.Println(multiply("123", "456"))
}

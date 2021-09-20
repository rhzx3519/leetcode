package main

import "fmt"

/**
https://leetcode-cn.com/problems/maximum-binary-string-after-change/
 */
func maximumBinaryString(binary string) string {
	n := len(binary)
	var i int
	for i < n && binary[i] == '1' {
		i++
	}
	if i == n {
		return binary
	}

	var count int
	for ; i < n; i++ {
		if binary[i] == '1' {
			count++
		}
	}
	bytes := []byte{}
	for i := 0; i < n - count - 1; i++ {
		bytes = append(bytes, '1')
	}
	bytes = append(bytes, '0')
	for i := 0; i < count; i++ {
		bytes = append(bytes, '1')
	}

	return string(bytes)
}

func main() {
	fmt.Println(maximumBinaryString("000110"))
	fmt.Println(maximumBinaryString("01"))
}



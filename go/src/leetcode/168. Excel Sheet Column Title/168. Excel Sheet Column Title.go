package main

import "fmt"

func convertToTitle(c int) string {
	N := 26
	bytes := []byte{}
	for c > 0 {
		c--
		bytes = append(bytes, byte(c%N) + 'A')
		c /= N
	}
	//fmt.Println(bytes)
	reverse(bytes)
	return string(bytes)
}

func reverse(bytes []byte) {
	l, r := 0, len(bytes)-1
	for l < r {
		bytes[l], bytes[r] = bytes[r], bytes[l]
		l++
		r--
	}
}

func main() {
	fmt.Println(convertToTitle(1))
	fmt.Println(convertToTitle(28))
	fmt.Println(convertToTitle(701))
}

package main

import "fmt"

/**
求序列{A, B, C, ... , Z, AA, AB, AC, ... , AZ, BA, BB, ... , AAA, AAB, ...} 的第N项
N转为26进制
 */
const N = 26
func boo(n int) string {
	bytes := []byte{}
	for n != 0 {
		n--
		bytes = append(bytes, byte('A' + n % N))
		n /= N
	}
	l, r := 0, len(bytes) - 1
	for l < r {
		bytes[l], bytes[r] = bytes[r], bytes[l]
		l++
		r--
	}
	fmt.Println(string(bytes))
	return string(bytes)
}

func main() {
	boo(4)					// D
	boo(26)					// Z
	boo(27)					// AA
	boo(52)					// AZ
	boo(53)					// BA
	boo(79)					// CA
	boo(26*26 + 26 + 1)		// AAA
}

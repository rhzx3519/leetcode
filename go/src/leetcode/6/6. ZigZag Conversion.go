package main

import (
	"fmt"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := make([][]byte, numRows)
	var k, d = 0, -1
	for i := range s {
		rows[k] = append(rows[k], s[i])
		if k == 0 {
			d = 1
		} else if k == numRows-1 {
			d = -1
		}
		k += d
	}
	fmt.Println(rows)
	bytes := []byte{}
	for i := range rows {
		bytes = append(bytes, rows[i]...)
	}
	return string(bytes)
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 4))
	fmt.Println(convert("A", 1))
}

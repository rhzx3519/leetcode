package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	a := make([][]byte, numRows)
	for i, _ := range a {
		a[i] = []byte{}
	}
	n := len(s)
	for i, k, f := 0, 0, 1; i < n;  i, k = i+1, k+f {
		a[k] = append(a[k], s[i])
		if k == 0 {
			f = 1
		} else if k == numRows-1 {
			f = -1
		}
	}
	//str := fmt.Sprintf("%s", a)
	b := []string{}
	for _, tmp := range a {
		b = append(b, string(tmp))
	}

	str := strings.ReplaceAll(strings.Join(b, ""), " ", "")
	return str
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 4))
	fmt.Println(convert("A", 1))
}

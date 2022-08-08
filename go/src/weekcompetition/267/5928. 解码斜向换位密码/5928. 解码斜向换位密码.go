/**
@author ZhengHao Lou
Date    2021/11/14
*/
package main

import (
	"fmt"
	"strings"
)

func decodeCiphertext(encodedText string, rows int) string {
	n := len(encodedText)
	cols := n / rows
	matrix := make([][]byte, rows)
	for i := range matrix {
		matrix[i] = make([]byte, cols)
	}
	var k int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			matrix[i][j] = encodedText[k]
			k++
		}
	}
	//fmt.Println(matrix)

	bytes := []byte{}
	for c := 0; c < cols; c++ {
		for i, j := 0, c; i < rows && j < cols; i, j = i+1, j+1 {
			bytes = append(bytes, matrix[i][j])
		}
	}

	return strings.TrimRight(string(bytes), " ")
}

func main() {
	fmt.Println(decodeCiphertext("ch   ie   pr", 3))
	fmt.Println(decodeCiphertext("iveo    eed   l te   olc", 4))
	fmt.Println(decodeCiphertext("coding", 1))
	fmt.Println(decodeCiphertext(" b  ac", 2))
}
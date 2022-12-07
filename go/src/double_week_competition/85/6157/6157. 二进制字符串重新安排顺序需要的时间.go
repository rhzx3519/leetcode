/**
@author ZhengHao Lou
Date    2022/08/20
*/
package main

import "fmt"

func secondsToRemoveOccurrences(s string) int {
	n := len(s)
	var c int
	bytes := []byte(s)
	for {
		var f bool
		for i := 0; i < n-1; {
			if string(bytes[i:i+2]) == "01" {
				f = true
				bytes[i] = '1'
				bytes[i+1] = '0'
				i += 2
			} else {
				i++
			}
		}
		if !f {
			return c
		}
		c++
	}

	return 0
}

func main() {
	fmt.Println(secondsToRemoveOccurrences("0110101")) // 1010101 -> 1100101 -> 1101001
	fmt.Println(secondsToRemoveOccurrences("11100"))
	fmt.Println(secondsToRemoveOccurrences("001011")) // 010101 -> 101010 ->
}

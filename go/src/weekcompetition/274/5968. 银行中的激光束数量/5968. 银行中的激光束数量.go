/**
@author ZhengHao Lou
Date    2022/01/02
*/
package main

import "fmt"

func numberOfBeams(bank []string) int {
	var total int
	var last int
	for i := range bank {
		var cur int
		for j := range bank[i] {
			if bank[i][j] == '1' {
				cur++
			}
		}

		if cur == 0 {
			continue
		}

		total += last * cur
		last = cur
	}

	return total
}

func main() {
	fmt.Println(numberOfBeams([]string{"011001","000000","010100","001000"}))
	fmt.Println(numberOfBeams([]string{"000","111","000"}))
	//fmt.Println(numberOfBeams([]string{}))
}

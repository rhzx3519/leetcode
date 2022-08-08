/**
@author ZhengHao Lou
Date    2022/01/22
*/
package main

import (
	"fmt"
)

func numberOfArrays(differences []int, lower int, upper int) int {
	//n := len(differences)
	var mi, mx int = 1e6, -1e6
	var t int
	for _, d := range differences {
		t += d
		if t < mi {
			mi = t
		}
		if t > mx {
			mx = t
		}
	}
	fmt.Println(mi, mx)
	var count int
	for i := lower; i <= upper; i++ {
		if i+mi >= lower && i+mx <= upper {
			count++
		}
	}
	fmt.Println(count)
	return count
}

func main() {
	numberOfArrays([]int{1, -3, 4}, 1, 6)
	numberOfArrays([]int{3, -4, 5, 1, -2}, -4, 5)
	numberOfArrays([]int{4, -7, 2}, 3, 6)
}

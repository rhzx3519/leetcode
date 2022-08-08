/**
@author ZhengHao Lou
Date    2021/12/04
*/
package main

import "fmt"

func pow(x float64, n int) float64 {
	var res float64 = 1.0
	var x_contribute = x
	for n > 0 {
		if n&1 == 1 {
			res *= x_contribute
		}
		x_contribute *= x_contribute
		n >>= 1
	}
	return res
}

func main() {
	fmt.Println(pow(2, 10))
}
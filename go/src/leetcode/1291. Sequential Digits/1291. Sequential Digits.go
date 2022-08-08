/**
@author ZhengHao Lou
Date    2021/12/03
*/
package main

import "fmt"

const N = 10
func sequentialDigits(low int, high int) []int {
	var l, d int
	var tmp = low
	for tmp != 0 {
		d = tmp % 10
		tmp /= 10
		l++
	}

	var ans = []int{}
	for {
		if 9-d+1 < l {
			d = 1
			l++
		}
		var c int
		for i := 0; i < l; i++ {
			c = c*10 + (d + i)
		}
		d++
		if c > high {
			break
		}
		if c >= low {
			ans = append(ans, c)
		}
	}

	fmt.Println(ans)
	return ans
}

func main() {
	//sequentialDigits(1000, 13000)
	sequentialDigits(58, 155)
}

/**
@author ZhengHao Lou
Date    2022/10/16
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/number-of-valid-clock-times/
*/
func countTime(time string) (c int) {
	var c1, c2 = 1, 1
	if time[0] == '?' && time[1] == '?' {
		c1 = 24
	} else if time[0] == '?' {
		if int(time[1]-'0') <= 3 {
			c1 = 3
		} else {
			c1 = 2
		}
	} else if time[1] == '?' {
		if int(time[0]-'0') == 2 {
			c1 = 4
		} else {
			c1 = 10
		}
	}

	if time[3] == '?' && time[4] == '?' {
		c2 = 60
	} else if time[3] == '?' {
		c2 = 6
	} else if time[4] == '?' {
		c2 = 10
	}

	return c1 * c2
}

func main() {
	fmt.Println(countTime("?5:00"))
	fmt.Println(countTime("0?:0?"))
	fmt.Println(countTime("??:??"))
}

/**
@author ZhengHao Lou
Date    2022/06/20
*/
package main

import "fmt"

/**
 */
func minimumNumbers(num int, k int) int {
	if num == 0 {
		return 0
	}
	for i := 1; i <= num && num-i*k >= 0; i++ {
		if (num-i*k)%10 == 0 {
			fmt.Println(i)
			return i
		}
	}
	return -1
}

func main() {
	minimumNumbers(58, 9)
	minimumNumbers(37, 2)
	minimumNumbers(0, 7)
	minimumNumbers(1, 1)
	minimumNumbers(4, 0)
}

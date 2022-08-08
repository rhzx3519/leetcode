/**
@author ZhengHao Lou
Date    2021/10/18
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/number-complement/
 */
func findComplement(num int) int {
	var x, l int
	for num != 0 {
		if num & 1 == 0 {
			x |= 1 << l
		}
		num >>= 1
		l++
	}
	return x
}

func main() {
	fmt.Println(findComplement(1))
}

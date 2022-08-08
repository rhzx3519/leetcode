/**
@author ZhengHao Lou
Date    2021/12/31
*/
package main

import "fmt"

func checkPerfectNumber(num int) bool {
	var count = 1
	for i := 2; i <= num/i; i++ {
		if num%i == 0 {
			fmt.Println(i)
			count += i + num/i
		}
	}
	fmt.Println(count)
	return count == num
}

func main() {
	//fmt.Println(checkPerfectNumber(28))
	//fmt.Println(checkPerfectNumber(6))
	//fmt.Println(checkPerfectNumber(496))
	fmt.Println(checkPerfectNumber(8128))
	//fmt.Println(checkPerfectNumber(2))
}
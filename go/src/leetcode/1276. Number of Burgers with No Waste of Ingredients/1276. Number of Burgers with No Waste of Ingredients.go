/**
@author ZhengHao Lou
Date    2021/12/02
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/number-of-burgers-with-no-waste-of-ingredients/
 */
func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
	for i := 0; i <= cheeseSlices; i++ {
		j := cheeseSlices - i
		if 4*i + 2*j == tomatoSlices {
			return []int{i, j}
		}
	}
	return []int{}
}

func main() {
	fmt.Println(numOfBurgers(16, 7))
	fmt.Println(numOfBurgers(17, 4))
	fmt.Println(numOfBurgers(4, 17))
	fmt.Println(numOfBurgers(0, 0))
	fmt.Println(numOfBurgers(2, 1))
}

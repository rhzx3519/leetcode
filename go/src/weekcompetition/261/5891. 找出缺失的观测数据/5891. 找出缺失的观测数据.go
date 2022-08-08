/**
@author ZhengHao Lou
Date    2021/10/03
*/
package main

import "fmt"

func missingRolls(rolls []int, mean int, n int) []int {
	var s int
	m := len(rolls)
	for i := range rolls {
		s += rolls[i]
	}
	total := mean * (m + n)
	t := total - s
	if t < n {
		return []int{}
	}
	x := t / n

	y := t % n
	if x > 6 || (x == 6 && y > 0) {
		return []int{}
	}
	ans := []int{}
	for i := 0; i < n; i++ {
		tmp := x
		if y > 0 {
			y--
			tmp += 1
		}
		ans = append(ans, tmp)
	}

	return ans
}

func main() {
	//fmt.Println(missingRolls([]int{3,2,4,3}, 4, 2))
	//fmt.Println(missingRolls([]int{1,5,6}, 3, 4))
	//fmt.Println(missingRolls([]int{1,2,3,4}, 6, 4))
	//fmt.Println(missingRolls([]int{1}, 3, 1))
	fmt.Println(missingRolls([]int{3,5,3}, 5, 3))
}

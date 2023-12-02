package main

import "fmt"

/*
*
https://leetcode.cn/problems/car-pooling/?envType=daily-question&envId=2023-12-02
*/
func carPooling(trips [][]int, capacity int) bool {
	f := make([]int, 1001)
	for i := range trips {
		num, from, to := trips[i][0], trips[i][1], trips[i][2]
		f[from] += num
		f[to] -= num
	}
	var c int
	for i := range f {
		c += f[i]
		if c > capacity {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(carPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 4))
	fmt.Println(carPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 5))
}

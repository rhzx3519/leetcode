/**
@author ZhengHao Lou
Date    2021/11/22
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/watering-plants/
思路：模拟
time: O(n)
space: O(1)
 */
func wateringPlants(plants []int, capacity int) int {
	var i, step int
	var cap = capacity
	n := len(plants)

	for i < n {
		for i < n && cap - plants[i] >= 0{
			cap -= plants[i]
			step++
			i++
		}
		if i < n {
			cap = capacity
			step += i<<1
		}
	}

	fmt.Println(step)
	return step
}

func main() {
	wateringPlants([]int{2,2,3,3}, 5)
	wateringPlants([]int{1,1,1,4,2,3}, 4)
	wateringPlants([]int{7,7,7,7,7,7,7}, 8)
}

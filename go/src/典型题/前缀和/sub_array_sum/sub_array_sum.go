package main

import "fmt"

/**
 给定一个int数组，求所有和等于目标值的非空子数组数量。
 子数组是指下标连续的部分。
 */
func subArraySum(arr []int, target int) int {
	mapper := make(map[int]int)
	mapper[0] = 1
	var count, pre int

	for _, num := range arr {
		pre += num
		if n, ok := mapper[pre - target]; ok {
			count += n
		}
		mapper[pre]++
	}
	return count
}

func main() {
	fmt.Println(subArraySum([]int{4,2,2,4,0}, 4))
	fmt.Println(subArraySum([]int{1,3,1,4,5}, 4))
}


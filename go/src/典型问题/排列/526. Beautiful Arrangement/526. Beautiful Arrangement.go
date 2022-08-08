package main

import "fmt"

/**
https://leetcode-cn.com/problems/beautiful-arrangement/

 */
func countArrangement(n int) int {
	mapper := make(map[int]int)
	for i := 1; i <= n; i++ {
		mapper[i]++
	}

	var count int
	var backtrace func([]int)
	backtrace = func(arr []int) {
		if len(arr) == n {
			count++
			return
		}

		keys := []int{}
		for num, _ := range mapper {
			keys = append(keys, num)
		}

		i := len(arr) + 1
		for _, num := range keys {
			if num%i == 0 || i%num == 0 {
				delete(mapper, num)
				arr = append(arr, num)

				backtrace(arr)

				arr = arr[:len(arr)-1]
				mapper[num]++
			}
		}
	}

	backtrace([]int{})
	return count
}

func main() {
	fmt.Println(countArrangement(2))
	fmt.Println(countArrangement(4))	// 8
}

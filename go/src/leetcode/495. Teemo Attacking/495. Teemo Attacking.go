/**
@author ZhengHao Lou
Date    2021/11/10
*/
package main

import "fmt"

func findPoisonedDuration(timeSeries []int, duration int) int {
	var total int
	n := len(timeSeries)
	for i := 0; i < n - 1; i++ {
		if timeSeries[i] + duration >= timeSeries[i+1] {
			total += timeSeries[i+1] - timeSeries[i]
		} else {
			total += duration
		}
	}
	total += duration
	return total
}

func main() {
	fmt.Println(findPoisonedDuration([]int{1,4}, 2))
	fmt.Println(findPoisonedDuration([]int{1,2}, 2))
}
package main

import "fmt"

func addRungs(rungs []int, dist int) int {
	var count, last int
	n := len(rungs)
	for i := 0; i < n; {
		rung := rungs[i]
		if last + dist < rung {
			x := (rung - last - 1) / dist
			last += x * dist
			count += x
		} else {
			last = rung
			i++
		}
	}
	fmt.Println(count)
	return count
}

func main() {
	addRungs([]int{4,8,12,16}, 3)
//	addRungs([]int{3}, 1)
//	addRungs([]int{1,3,5,10}, 2)
//	addRungs([]int{3,6,8,10}, 3)
//	addRungs([]int{3,4,6,7}, 2)
//	addRungs([]int{5}, 10)
}
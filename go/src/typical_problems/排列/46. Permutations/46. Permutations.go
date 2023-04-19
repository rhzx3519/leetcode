package main

import "fmt"

func permute(nums []int) [][]int {
	perms := [][]int{}

	n := len(nums)
	var backtrace func([]int, map[int]bool)
	backtrace = func(a []int, mapper map[int]bool) {
		if len(a) == n {
			perms = append(perms, append([]int{}, a...))
			return
		}

		for _, num := range nums {
			if _, ok := mapper[num]; ok {
				continue
			}
			a = append(a, num)
			mapper[num] = true
			backtrace(a, mapper)
			delete(mapper, num)
			a = a[:len(a)-1]
		}
	}

	backtrace([]int{}, make(map[int]bool))
	return perms
}

func main() {
	fmt.Println(permute([]int{1,2,3}))
}
package main

import "fmt"

func rearrangeArray(nums []int) []int {
	mapper := make(map[int]bool)
	for _, num := range nums {
		mapper[num] = true
	}
	n := len(nums)

	ans := []int{}

	var backtrace func([]int) bool
	backtrace = func(arr []int) bool {
		if len(arr) == n {
			ans = append(ans, arr...)
			return true
		}

		for num, _ := range mapper {
			l := len(arr)
			if l >= 2 && arr[l-1]<<1 == arr[l-2] + num {
				continue
			}
			delete(mapper, num)
			arr = append(arr, num)
			if backtrace(arr) {
				return true
			}
			arr = arr[:len(arr)-1]
			mapper[num] = true
		}
		return false
	}

	backtrace([]int{})
	fmt.Println(ans)
	return ans
}

func main() {
	rearrangeArray([]int{1,2,3,4,5})
	rearrangeArray([]int{6,2,0,9,7})
}

/**
@author ZhengHao Lou
Date    2021/12/26
*/
package main

import "fmt"

func getDistances(arr []int) []int64 {
	n := len(arr)
	var left = make([]int64, n)
	counter := make(map[int]int)
	mapper := make(map[int]int)
	for i, num := range arr {
		if j, ok := mapper[num]; ok {
			left[i] += int64((i - j)*counter[num]) + left[j]
		}
		mapper[num] = i
		counter[num]++
	}


	var right = make([]int64, n)
	counter = make(map[int]int)
	mapper = make(map[int]int)
	for i := len(arr) - 1; i >= 0; i-- {
		num := arr[i]
		if j, ok := mapper[num]; ok {
			right[i] += int64((j - i)*counter[num]) + right[j]
		}
		mapper[num] = i
		counter[num]++
	}


	//fmt.Println(counter)
	//fmt.Println(left)
	//fmt.Println(right)
	var ans = make([]int64, n)
	for i := range left {
		ans[i] = left[i] + right[i]
	}
	fmt.Println(ans)
	return ans
}


func main() {
	getDistances([]int{2,1,3,1,2,3,3})
	getDistances([]int{10,5,10,10})
}

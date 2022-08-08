package main

import "fmt"

func canBeIncreasing(nums []int) bool {
	n := len(nums)
	var i1, i2 = -1, -1
	for i := 1; i < n; i++ {
		if nums[i] <= nums[i-1] {
			i1 = i-1
			i2 = i
			break
		}
	}
	if i1 == -1 && i2 == -1 {
		return true
	}
	tmp := []int{}
	tmp = append(tmp, nums[:i1]...)
	tmp = append(tmp, nums[i1+1:]...)
	if inc(tmp) {
		return true
	}
	tmp = []int{}
	tmp = append(tmp, nums[:i2]...)
	tmp = append(tmp, nums[i2+1:]...)
	if inc(tmp) {
		return true
	}
	return false
}

func inc(a []int) bool {
	n := len(a)
	for i := 1; i < n; i++ {
		if a[i] <= a[i-1] {
			return false
		}
	}
	return true
}

func main() {
	//fmt.Println(canBeIncreasing([]int{1,2,10,5,7}))
	//fmt.Println(canBeIncreasing([]int{2,3,1,2}))
	fmt.Println(canBeIncreasing([]int{1,1,1}))
	//fmt.Println(canBeIncreasing([]int{1,2,3}))
}


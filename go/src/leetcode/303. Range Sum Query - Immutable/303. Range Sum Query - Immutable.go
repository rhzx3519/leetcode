package main

import "fmt"

type NumArray struct {
	pre []int
}


func Constructor(nums []int) NumArray {
	obj := NumArray{make([]int, len(nums)+1)}
	for i := 1; i < len(nums)+1; i++ {
		obj.pre[i] = obj.pre[i-1] + nums[i-1]
	}
	return obj
}


func (this *NumArray) SumRange(i int, j int) int {
	return this.pre[j+1] - this.pre[i]
}

func main() {
	su := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(su.SumRange(0, 2))
	fmt.Println(su.SumRange(2,5))
	fmt.Println(su.SumRange(0,5))
}


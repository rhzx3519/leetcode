package main

import (
	"fmt"
	"sort"
)

type MKAverage struct {
	m, k int
	stream []int
}


func Constructor(m int, k int) MKAverage {
	return MKAverage{
		m: m,
		k: k,
		stream: make([]int, 0),
	}
}


func (this *MKAverage) AddElement(num int)  {
	idx := searchInsert(this.stream, num)
	insert(&this.stream, idx, num)
}


func (this *MKAverage) CalculateMKAverage() int {
	if len(this.stream) < this.m {
		return -1
	}
	tmp := make([]int, this.m)
	copy(tmp, this.stream[len(this.stream)-this.m:])
	sort.Ints(tmp)
	s := 0
	for i := this.k; i < this.m-this.k; i++ {
		s += tmp[i]
	}
	fmt.Println(s / (this.m-2*this.k))
	return s / (this.m-2*this.k)
}

func searchInsert(nums []int, target int) int {
	n := len(nums)
	ans := n	// ans初始化为n，省略边界判断，有一种情况是target > nums[-1]
	var mid int
	for l, r := 0, n-1; l <= r; {
		mid = (r - l)>>1 + l
		if nums[mid] >= target {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ans
}

func insert(a *[]int, idx, x int) {
	if len(*a) == cap(*a) {
		*a= append(*a, x)
	}
	tmp := *a
	copy(tmp[idx+1:], tmp[idx:])
	tmp[idx] = x
}


/**
 * Your MKAverage object will be instantiated and called as such:
 * obj := Constructor(m, k);
 * obj.AddElement(num);
 * param_2 := obj.CalculateMKAverage();
 */

func main() {
	m, k := 3, 1
	obj := Constructor(m, k);
	obj.AddElement(3)
	obj.AddElement(1)
	obj.CalculateMKAverage()
	obj.AddElement(10)
	obj.CalculateMKAverage()
	obj.AddElement(5)
	obj.AddElement(5)
	obj.AddElement(5)
	obj.CalculateMKAverage()

	var a = []int{0,1}
	fmt.Println(a, len(a), cap(a))
	insert(&a, 0, -1)
	fmt.Println(a, len(a), cap(a))
}
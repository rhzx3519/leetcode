/**
@author ZhengHao Lou
Date    2023/01/18
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode.cn/problems/finding-mk-average/
*/
type MKAverage struct {
	m, k, s        int
	nums           []int
	low, mid, high []int
}

func Constructor(m int, k int) MKAverage {
	return MKAverage{m: m, k: k}
}

func (this *MKAverage) AddElement(num int) {
	if len(this.low) == 0 || num <= this.low[len(this.low)-1] {
		insert(&this.low, num)
	} else if len(this.high) == 0 || num >= this.high[0] {
		insert(&this.high, num)
	} else {
		this.s += num
		insert(&this.mid, num)
	}
	this.nums = append(this.nums, num)
	if len(this.nums) > this.m {
		v := this.nums[0]
		this.nums = this.nums[1:]
		if j := search(&this.low, v); j != -1 {
			delete(&this.low, j)
		} else if j := search(&this.high, v); j != -1 {
			delete(&this.high, j)
		} else if j := search(&this.mid, v); j != -1 {
			this.s -= this.mid[j]
			delete(&this.mid, j)
		}
	}
	
	for len(this.low) > this.k {
		v := this.low[len(this.low)-1]
		this.s += v
		insertIn(&this.mid, 0, v)
		this.low = this.low[:len(this.low)-1]
	}
	
	for len(this.high) > this.k {
		v := this.high[0]
		this.s += v
		this.mid = append(this.mid, v)
		this.high = this.high[1:]
	}
	
	for len(this.low) < this.k && len(this.mid) != 0 {
		this.low = append(this.low, this.mid[0])
		this.s -= this.mid[0]
		this.mid = this.mid[1:]
	}
	for len(this.high) < this.k && len(this.mid) != 0 {
		v := this.mid[len(this.mid)-1]
		insertIn(&this.high, 0, v)
		
		this.s -= v
		this.mid = this.mid[:len(this.mid)-1]
	}
}

func (this *MKAverage) CalculateMKAverage() int {
	if len(this.nums) < this.m {
		return -1
	}
	return this.s / (this.m - this.k<<1)
}

func insert(nums *[]int, x int) {
	j := sort.Search(len(*nums), func(i int) bool {
		return (*nums)[i] >= x
	})
	*nums = append(*nums, 0)
	copy((*nums)[j+1:], (*nums)[j:])
	(*nums)[j] = x
}

func insertIn(nums *[]int, j, x int) {
	*nums = append(*nums, 0)
	copy((*nums)[j+1:], (*nums)[j:])
	(*nums)[j] = x
}

func delete(nums *[]int, j int) {
	copy((*nums)[j:], (*nums)[j+1:])
	*nums = (*nums)[:len(*nums)-1]
}

func search(nums *[]int, x int) int {
	j := sort.Search(len(*nums), func(i int) bool {
		return (*nums)[i] >= x
	})
	if j != len(*nums) && (*nums)[j] == x {
		return j
	}
	return -1
}

func main() {
	mk := Constructor(3, 1)
	mk.AddElement(17612)
	mk.AddElement(74607)
	fmt.Println(mk.CalculateMKAverage())
	mk.AddElement(8272)
	mk.AddElement(33433)
	fmt.Println(mk.CalculateMKAverage())
	mk.AddElement(15456)
	mk.AddElement(64938)
	fmt.Println(mk.CalculateMKAverage())
	mk.AddElement(99741)
	fmt.Println(mk.CalculateMKAverage())
}

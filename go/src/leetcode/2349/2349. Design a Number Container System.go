/**
@author ZhengHao Lou
Date    2022/07/25
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/design-a-number-container-system/
*/
type NumberContainers struct {
	indice  map[int]int
	numbers map[int][]int
}

func Constructor() NumberContainers {
	return NumberContainers{
		indice:  map[int]int{},
		numbers: map[int][]int{},
	}
}

func (this *NumberContainers) Change(index int, number int) {
	if x, ok := this.indice[index]; ok {
		i := lowerBound(index, this.numbers[x])
		copy(this.numbers[x][i:], this.numbers[x][i+1:])
		this.numbers[x] = this.numbers[x][:len(this.numbers[x])-1]
	}
	this.indice[index] = number
	i := lowerBound(index, this.numbers[number])
	this.numbers[number] = append(this.numbers[number], 0)
	copy(this.numbers[number][i+1:], this.numbers[number][i:])
	this.numbers[number][i] = index
}

func (this *NumberContainers) Find(number int) int {
	if len(this.numbers[number]) == 0 {
		return -1
	}
	fmt.Println(this.numbers[number])
	return this.numbers[number][0]
}

func lowerBound(x int, nums []int) int {
	l, r := 0, len(nums)
	for l < r {
		m := l + (r-l)>>1
		if nums[m] >= x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func main() {
	obj := Constructor()
	obj.Change(2, 10)
	obj.Change(1, 10)
	obj.Change(3, 10)
	obj.Change(5, 10)
	fmt.Println(obj.Find(10))
	obj.Change(1, 20)
	fmt.Println(obj.Find(10))
}

/**
 * Your NumberContainers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Change(index,number);
 * param_2 := obj.Find(number);
 */

package main

import "fmt"

type MedianFinder struct {
	arr []int
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		arr: []int{},
	}
}


func (this *MedianFinder) AddNum(num int)  {
	i := lowerBound(this.arr, num)
	this.arr = append(this.arr, 0)
	copy(this.arr[i+1:], this.arr[i:])
	this.arr[i] = num
}


func (this *MedianFinder) FindMedian() float64 {
	n := len(this.arr)
	return float64(this.arr[(n)>>1] + this.arr[(n-1)>>1]) * 0.5
}

func lowerBound(arr []int, target int) int {
	l, r := 0, len(arr)
	var m int
	for l < r {
		m = l + (r-l)>>1
		if arr[m] >= target {
			r = m
		} else {
			l++
		}
	}
	return l
}

func main() {
	obj := Constructor()
	obj.AddNum(1)
	obj.AddNum(3)
	fmt.Println(obj.FindMedian())
	obj.AddNum(-1)
	obj.AddNum(2)
	fmt.Println(obj.arr)
	fmt.Println(obj.FindMedian())
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

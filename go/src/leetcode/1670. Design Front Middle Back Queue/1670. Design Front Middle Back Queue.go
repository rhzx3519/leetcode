/**
@author ZhengHao Lou
Date    2021/11/05
*/
package main

type FrontMiddleBackQueue struct {
	a []int
}


func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{
		a: []int{},
	}
}


func (this *FrontMiddleBackQueue) PushFront(val int)  {
	this.a = append(this.a, 0)
	copy(this.a[1:], this.a)
	this.a[0] = val
}


func (this *FrontMiddleBackQueue) PushMiddle(val int)  {
	var m = len(this.a) >> 1
	this.a = append(this.a, 0)
	copy(this.a[m+1:], this.a[m:])
	this.a[m] = val
}


func (this *FrontMiddleBackQueue) PushBack(val int)  {
	this.a = append(this.a, 0)
	this.a[len(this.a) - 1] = val
}


func (this *FrontMiddleBackQueue) PopFront() int {
	if len(this.a) == 0 {
		return -1
	}
	var val = this.a[0]
	copy(this.a, this.a[1:])
	this.a = this.a[:len(this.a) - 1]
	return val
}


func (this *FrontMiddleBackQueue) PopMiddle() int {
	if len(this.a) == 0 {
		return -1
	}
	var m = (len(this.a) - 1) >> 1
	var val = this.a[m]
	copy(this.a[m:], this.a[m+1:])
	this.a = this.a[:len(this.a) - 1]
	return val
}


func (this *FrontMiddleBackQueue) PopBack() int {
	if len(this.a) == 0 {
		return -1
	}
	var val = this.a[len(this.a) - 1]
	this.a = this.a[:len(this.a) - 1]
	return val
}


/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */
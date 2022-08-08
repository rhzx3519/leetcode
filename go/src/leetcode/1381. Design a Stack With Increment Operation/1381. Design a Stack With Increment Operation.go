/**
@author ZhengHao Lou
Date    2021/11/23
*/
package main

/**
https://leetcode-cn.com/problems/design-a-stack-with-increment-operation/
 */
type CustomStack struct {
	st []int
	capacity int
}


func Constructor(maxSize int) CustomStack {
	return CustomStack{
		st: []int{},
		capacity: maxSize,
	}
}


func (this *CustomStack) Push(x int)  {
	if len(this.st) >= this.capacity {
		return
	}
	this.st = append(this.st, x)
}


func (this *CustomStack) Pop() int {
	if len(this.st) > 0 {
		v := this.st[len(this.st) - 1]
		this.st = this.st[:len(this.st) - 1]
		return v
	}
	return -1
}


func (this *CustomStack) Increment(k int, val int)  {
	k = min(k, len(this.st))
	for i := range this.st[:k] {
		this.st[i] += val
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}


/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */

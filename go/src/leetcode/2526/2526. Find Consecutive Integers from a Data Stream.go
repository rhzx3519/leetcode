/**
@author ZhengHao Lou
Date    2023/01/08
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/find-consecutive-integers-from-a-data-stream/
*/
type DataStream struct {
	value, k int
	last     int
}

func Constructor(value int, k int) DataStream {
	return DataStream{value: value, k: k}
}

func (this *DataStream) Consec(num int) bool {
	if num != this.value {
		this.last = 0
	} else {
		this.last++
	}
	return this.last >= this.k
}

func main() {
	obj := Constructor(4, 3)
	fmt.Println(obj.Consec(4))
	fmt.Println(obj.Consec(4))
	fmt.Println(obj.Consec(4))
	fmt.Println(obj.Consec(3))
	fmt.Println(obj.Consec(4))
	fmt.Println(obj.Consec(4))
	fmt.Println(obj.Consec(3))
}

/**
 * Your DataStream object will be instantiated and called as such:
 * obj := Constructor(value, k);
 * param_1 := obj.Consec(num);
 */

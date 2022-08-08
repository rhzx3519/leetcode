/**
@author ZhengHao Lou
Date    2022/01/12
*/
package main

/**
https://leetcode-cn.com/problems/maximum-frequency-stack/
 */
type FreqStack struct {
	freq map[int]int	// 记录每个元素的频率
	group map[int][]int	// freq: vals
	maxFreq int			// 记录 当前最大频率
}


func Constructor() FreqStack {
	return FreqStack{
		freq: map[int]int{},
		group: map[int][]int{},
	}
}


func (this *FreqStack) Push(val int)  {
	this.freq[val]++
	var f = this.freq[val]
	if f > this.maxFreq {
		this.maxFreq = f
	}
	if _, ok := this.group[f]; !ok {
		this.group[f] = []int{}
	}
	this.group[f] = append(this.group[f], val)
}


func (this *FreqStack) Pop() int {
	x := this.group[this.maxFreq][len(this.group[this.maxFreq]) - 1]
	this.group[this.maxFreq] = this.group[this.maxFreq][:len(this.group[this.maxFreq]) - 1]
	this.freq[x]--
	if len(this.group[this.maxFreq]) == 0 {
		this.maxFreq--
	}
	return x
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */

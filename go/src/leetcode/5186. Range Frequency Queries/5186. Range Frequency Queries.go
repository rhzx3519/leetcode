/**
@author ZhengHao Lou
Date    2021/11/22
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/range-frequency-queries/
思路：区间查询
time: O(logn)
space: O(n)
 */
type RangeFreqQuery struct {
	tree *FenwickTree
}


func Constructor(arr []int) RangeFreqQuery {
	r := RangeFreqQuery{
		tree: NewFenwickTree(len(arr)),
	}
	for i, num := range arr {
		r.tree.Update(i + 1, num)
	}

	return r
}


func (this *RangeFreqQuery) Query(left int, right int, value int) int {
	if left == 0 {
		return this.tree.Freq(right + 1, value)
	}
	l, r := this.tree.Freq(left, value), this.tree.Freq(right + 1, value)
	return r - l
}

// --------------------------------------------------
type FenwickTree struct {
	freq []map[int]int
}

func (f *FenwickTree) Update(i, k int) {
	for i < len(f.freq) {
		f.freq[i][k] ++
		i += lowbit(i)
	}
}

// 返回[0,i]之间的元素和
func (f *FenwickTree) Freq(i int, value int) int {
	var s int
	for i > 0 {
		s += f.freq[i][value]
		i -= lowbit(i)
	}
	return s
}

func lowbit(x int) int {
	return x & (-x)
}

func NewFenwickTree(n int) *FenwickTree {
	freq := make([]map[int]int, n + 1)
	for i := range freq {
		freq[i] = make(map[int]int)
	}
	return &FenwickTree{
		freq: freq,
	}
}


/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,value);
 */

func main() {
	r := Constructor([]int{12, 33, 4, 56, 22, 2, 34, 33, 22, 12, 34, 56})
	fmt.Println(r.Query(1, 2, 4))
	fmt.Println(r.Query(0, 11, 33))
}
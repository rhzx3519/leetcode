/*
*
@author ZhengHao Lou
Date    2021/11/22
*/
package main

import "fmt"

// 树状数组的下标为[1, n]
type FenwickTree struct {
	sums []int
}

func (f *FenwickTree) Update(i, k int) {
	for i < len(f.sums) {
		f.sums[i] += k
		i += lowbit(i)
	}
}

// 返回[1,i]之间的元素和
func (f *FenwickTree) Range(i int) int {
	var s int
	for i > 0 {
		s += f.sums[i]
		i -= lowbit(i)
	}
	return s
}

func lowbit(x int) int {
	return x & (-x)
}

func NewFenwickTree(n int) *FenwickTree {
	return &FenwickTree{
		sums: make([]int, n+1),
	}
}

func main() {
	tree := NewFenwickTree(10)
	tree.Update(1, 5)
	fmt.Println(tree.Range(1)) // 5
	fmt.Println(tree.Range(0)) // 0
	fmt.Println(tree.Range(9)) // 5
}

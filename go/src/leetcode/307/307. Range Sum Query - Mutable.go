/**
@author ZhengHao Lou
Date    2022/04/04
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/range-sum-query-mutable/
*/
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

type NumArray struct {
	nums []int
	tree *FenwickTree
}

func Constructor(nums []int) NumArray {
	tree := NewFenwickTree(len(nums))
	for i := range nums {
		tree.Update(i+1, nums[i])
	}

	return NumArray{
		nums: nums,
		tree: tree,
	}
}

func (this *NumArray) Update(index int, val int) {
	this.tree.Update(index+1, val-this.nums[index])
	this.nums[index] = val
}

func (this *NumArray) SumRange(left int, right int) int {
	fmt.Println(this.tree.Range(right+1) - this.tree.Range(left))
	return this.tree.Range(right+1) - this.tree.Range(left)
}

func main() {
	obj := Constructor([]int{1, 3, 5})
	obj.SumRange(0, 2)
	obj.Update(1, 2)
	obj.SumRange(0, 2)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */


/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

package main

type NestedInteger struct {

}

func (this NestedInteger) IsInteger() bool {
	return false
}

func (this NestedInteger) GetInteger() int {
	return 0
}

func (this *NestedInteger) Add(elem NestedInteger) {}

func (this NestedInteger) GetList() []*NestedInteger {
	return nil
}

type NestedIterator struct {
	idx int
	que []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	iterator := new(NestedIterator)
	iterator.parse(nestedList)
	return iterator
}

func (this *NestedIterator) Next() int {
	next := this.que[this.idx]
	this.idx++
	return next
}

func (this *NestedIterator) HasNext() bool {
	return this.idx < len(this.que)
}

func (this *NestedIterator) parse(nestedList []*NestedInteger) {
	for _, nestedInteger := range nestedList {
		if nestedInteger.IsInteger() {
			this.que = append(this.que, nestedInteger.GetInteger())
		} else {
			this.parse(nestedInteger.GetList())
		}
	}
}
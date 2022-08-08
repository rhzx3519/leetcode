/**
@author ZhengHao Lou
Date    2022/04/15
*/
package main

import "unicode"

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */
func deserialize(s string) *NestedInteger {
	n := len(s)
	var i int

	var dfs func() *NestedInteger
	dfs = func() *NestedInteger {
		var ni = &NestedInteger{}
		if s[i] == '[' {
			i++
			for i < n && s[i] != ']' {
				ni.Add(*dfs())
				if s[i] == ',' {
					i++
				}
			}
			i++
		} else {
			var neg bool
			if s[i] == '-' {
				neg = true
				i++
			}
			var d int
			for i < n && unicode.IsDigit(rune(s[i])) {
				d = d*10 + int(s[i]-'0')
				i++
			}
			if neg {
				d = -d
			}
			ni.SetInteger(d)
		}
		return ni
	}

	return dfs()
}

type NestedInteger struct{}

func (n NestedInteger) IsInteger() bool           { return true }
func (n NestedInteger) GetInteger() int           { return 0 }
func (n *NestedInteger) SetInteger(value int)     {}
func (n *NestedInteger) Add(elem NestedInteger)   {}
func (n NestedInteger) GetList() []*NestedInteger { return nil }

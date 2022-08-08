/**
@author ZhengHao Lou
Date    2022/07/26
*/
package main

/**
https://leetcode.cn/problems/design-skiplist/
*/
const TotalLevelNum = 30

type SkiplistNode struct {
	Val    int
	Up     *SkiplistNode
	Bottom *SkiplistNode
	Next   *SkiplistNode
}

type Skiplist struct {
	Heads [TotalLevelNum]*SkiplistNode
}

func Constructor() Skiplist {

}

func (this *Skiplist) Search(target int) bool {

}

func (this *Skiplist) Add(num int) {

}

func (this *Skiplist) Erase(num int) bool {

}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */

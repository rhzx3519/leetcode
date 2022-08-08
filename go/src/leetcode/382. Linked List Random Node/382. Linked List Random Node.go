/**
@author ZhengHao Lou
Date    2022/01/16
*/
package main

import "math/rand"

/**
https://leetcode-cn.com/problems/linked-list-random-node/
蓄水池算法：当样本总量未知时，使抽取每个样本的概率相同的算法
当前遍历到n, 则使用概率1/n的概率来决定是否采样，这样遍历完整个样本后，每个样本被抽取的概率相同
 */
type ListNode struct {
	Val int
	Next *ListNode
}

type Solution struct {
	head *ListNode
}


func Constructor(head *ListNode) Solution {
	return Solution{head: head}
}


func (this *Solution) GetRandom() int {
	var ans int
	for i, p := 1, this.head; p != nil; i++ {
		if rand.Intn(i)	== 0 {
			ans = p.Val
		}
		p = p.Next
	}
	return ans
}


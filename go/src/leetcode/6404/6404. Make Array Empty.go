package main

import "sort"

/*
*
https://leetcode.cn/problems/make-array-empty/
*/
func countOperationsToEmptyArray(nums []int) int64 {
	n := len(nums)
	id := make([]int, n)
	for i := range id {
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool {
		return nums[id[i]] < nums[id[j]]
	})

	btree := make(BIT, n+1)
	ans := n
	pre := 1
	for k, i := range id {
		i++
		if i >= pre {
			ans += i - pre - btree.query(pre, i)
		} else {
			ans += i - pre + n - k + btree.query(i, pre-1)
		}
		btree.inc(i)
		pre = i
	}

	return int64(ans)
}

// 树状数组模板
type BIT []int

// 将下标 i 上的数加一
func (t BIT) inc(i int) {
	for ; i < len(t); i += i & -i {
		t[i]++
	}
}

// 返回闭区间 [1, i] 的元素和
func (t BIT) sum(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += t[i]
	}
	return
}

// 返回闭区间 [left, right] 的元素和
func (t BIT) query(left, right int) int {
	return t.sum(right) - t.sum(left-1)
}

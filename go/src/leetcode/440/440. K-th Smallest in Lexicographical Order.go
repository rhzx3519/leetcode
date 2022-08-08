/**
@author ZhengHao Lou
Date    2022/03/23
*/
package main

/**
https://leetcode-cn.com/problems/k-th-smallest-in-lexicographical-order/
*/
func findKthNumber(n int, k int) int {
	var p, prefix = 1, 1
	for p < k {
		count := getCount(prefix, n) // 获得当前前缀下所有子节点的和
		if p+count > k {             //第k个数在当前前缀下
			prefix *= 10
			p++ // 把指针指向了第一个子节点的位置，比如11乘10后变成110，指针从11指向了110
		} else if p+count <= k {
			prefix++
			p += count
		}
	}
	return prefix
}

// 确定指定前缀下所有子节点数
// prefix是前缀，n是上界
func getCount(prefix, n int) int {
	cur, next := prefix, prefix+1
	var count int
	for cur <= n {
		count += min(n+1, next) - cur // 下一个前缀的起点减去当前前缀的起点
		cur *= 10
		next *= 10
		// 如果说刚刚prefix是1，next是2，那么现在分别变成10和20
		// 1为前缀的子节点增加10个，十叉树增加一层, 变成了两层

		// 如果说现在prefix是10，next是20，那么现在分别变成100和200，
		// 1为前缀的子节点增加100个，十叉树又增加了一层，变成了三层
	}
	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

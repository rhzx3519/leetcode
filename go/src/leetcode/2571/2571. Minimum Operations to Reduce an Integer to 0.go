package main

import "fmt"

/*
*
https://leetcode.cn/problems/minimum-operations-to-reduce-an-integer-to-0/
思路：贪心
如果是连续的1，n = n + lb
如果是单个1，n = n - lb
lb = n&-n, 表示n最左边1的权重, e.g. 1010(10), lb = 10(2)
*/
func minOperations(n int) (tot int) {
	tot = 1
	for n&(n-1) != 0 {
		lb := n & -n
		if n&(lb<<1) > 0 { // consecutive 1
			n += lb
		} else {
			n -= lb
		}
		tot++
	}
	return
}

func main() {
	fmt.Println(minOperations(39))
	fmt.Println(minOperations(54))
}

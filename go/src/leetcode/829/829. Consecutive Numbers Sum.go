/**
@author ZhengHao Lou
Date    2022/06/03
*/
package main

func isKConsecutive(n, k int) bool {
	if k%2 == 1 {
		return n%k == 0
	}
	return n%k != 0 && 2*n%k == 0
}

func consecutiveNumbersSum(n int) (ans int) {
	for k := 1; k*(k+1) <= n*2; k++ {
		if isKConsecutive(n, k) {
			ans++
		}
	}
	return
}

//作者：LeetCode-Solution
//链接：https://leetcode.cn/problems/consecutive-numbers-sum/solution/lian-xu-zheng-shu-qiu-he-by-leetcode-sol-33hc/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

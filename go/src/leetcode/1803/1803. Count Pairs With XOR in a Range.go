/**
@author ZhengHao Lou
Date    2023/01/05
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/count-pairs-with-xor-in-a-range/
*/
func countPairs(nums []int, low, high int) (ans int) {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}
	for high++; high > 0; high >>= 1 {
		nxt := map[int]int{}
		for x, c := range cnt {
			ans += c * (high%2*cnt[(high-1)^x] - low%2*cnt[(low-1)^x])
			nxt[x>>1] += c
		}
		cnt = nxt
		low >>= 1
	}
	return ans / 2
}

//作者：灵茶山艾府
//链接：https://leetcode.cn/problems/count-pairs-with-xor-in-a-range/solutions/2045560/bu-hui-zi-dian-shu-zhi-yong-ha-xi-biao-y-p2pu/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	countPairs([]int{1, 4, 2, 7}, 2, 6)
	//countPairs([]int{9, 8, 4, 2, 1}, 5, 14)
	var nums = []int{1, 4, 2, 7}
	var cnt = map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}
	var tot int
	for x, c := range cnt {
		tot += c * cnt[x^3]
	}
	fmt.Println(tot, 4^7, 1^2)
}

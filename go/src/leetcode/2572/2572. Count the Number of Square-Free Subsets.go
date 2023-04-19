package main

/*
*
https://leetcode.cn/problems/count-the-number-of-square-free-subsets/
思路：状态压缩，01背包
*/
var (
	primes   = [...]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	nsq2mask = [31]int{} // nsq2mask为nsq对应的素数集合，nsq表示无平方因子的数字
)

func init() {
	for i := 2; i <= 30; i++ {
		for j, p := range primes {
			if i%p == 0 {
				if i%(p*p) == 0 {
					nsq2mask[i] = -1
					continue
				}
				nsq2mask[i] |= 1 << j // 添加j到集合中, 表示第j个素数在集合中
			}
		}
	}
}

func squareFreeSubsets(nums []int) int {
	const mod int = 1e9 + 7
	const m = 1 << len(primes)
	f := [m]int{1}
	for _, x := range nums {
		if mask := nsq2mask[x]; mask != -1 {
			for j := m - 1; j >= mask; j-- {
				if j|mask == j { // mask是j的子集
					f[j] = (f[j] + f[j^mask]) % mod
				}
			}
		}
	}

	var ans int
	for i := range f {
		ans = (ans + f[i]) % mod
	}

	return (ans - 1 + mod) % mod
}

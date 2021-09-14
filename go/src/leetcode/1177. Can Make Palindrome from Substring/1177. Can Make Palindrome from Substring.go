package main

import "fmt"

/**
思路：26个字母可以用一个32位整型表示，前缀异或和sum[r+1]-sum[l]表示[l,r]之前的各个字母奇偶个数。
 */
func canMakePaliQueries(s string, queries [][]int) []bool {
	n := len(s)
	sum := make([]uint32, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = (1<<int(s[i-1] - 'a')) ^ sum[i-1]
	}

	ans := []bool{}
	for _, q := range queries {
		l, r, k := q[0], q[1], q[2]
		ans = append(ans, isPal(sum, l, r, k))
	}

	return ans
}

const N = 26

func isPal(sum []uint32, l, r, k int) bool {
	t := sum[r+1] ^ sum[l]

	var odd int
	for i := 0; i < N; i++ {
		if t>>i&1 == 1 {
			odd++
		}
	}
	if (r - l + 1)&1 == 1 {
		odd--
	}
	k -= odd>>1
	return k >= 0
}

func main() {
	fmt.Println(canMakePaliQueries("abcda", [][]int{{3,3,0},{1,2,0},{0,3,1},{0,3,2},{0,4,1}}))
	fmt.Println(canMakePaliQueries("huxnu", [][]int{{0,4,1}}))
}

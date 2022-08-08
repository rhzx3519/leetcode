package main

import "fmt"

/**
字母出现次数全是偶数，或者只有一个字母出现了奇数次
 */
const N = 10

func wonderfulSubstrings(word string) int64 {
	var ans = int64(0)

	n := len(word)
	sum := make([]int, n+1)

	for i := 1; i <= n; i++ {
		k := word[i-1] - 'a'
		d := sum[i-1]>>k & 1 // k 出现d次

		d ^= 1
		x := 1<<k
		if d == 1 {
			sum[i] = sum[i-1] | x
		} else {
			sum[i] = sum[i-1] & ^(1<<k)
		}
	}

	mapper := make(map[int]int64)
	mapper[0] = int64(1)
	for i := range word {
		//k := word[i] - 'a'
		s := sum[i+1]
		if cnt, ok := mapper[s]; ok {
			ans += cnt
		}

		for c := 0; c < N; c++ {
			tmp := s
			if tmp>>c&1 == 1 { // 第c位为1
				tmp &= ^(1<<c)
			} else { // 第c位为0
				tmp |= 1<<c
			}
			if cnt, ok := mapper[tmp]; ok {
				ans += cnt
			}
		}

		mapper[s]++
	}

	fmt.Println(ans)
	return ans
}

func main() {
	wonderfulSubstrings("aba")
	wonderfulSubstrings("aabb")
	wonderfulSubstrings("he")
}
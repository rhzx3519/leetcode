package main

import "fmt"

func countPalindromicSubsequence(s string) int {
	n := len(s)
	pre := make([][]int, n+1)
	pre[0] = make([]int, 26)
	for i := 1; i <= n; i++ {
		pre[i] = make([]int, 0)
		pre[i] = append(pre[i], pre[i-1]...)
		pre[i][s[i-1] - 'a']++
	}
	//tmp := pre[4] ^ pre[1]
	//fmt.Printf("%b\n", tmp)
	count := make([]int, 26)

	mapper := make(map[byte]int)
	for i := range s {
		if _, ok := mapper[s[i]]; !ok {
			mapper[s[i]] = i
			continue
		}
		if j, ok := mapper[s[i]]; ok {
			tmp := inter(pre, j+1, i)
			count[int(s[i] - 'a')] = tmp
		}
	}

	var ans int
	for i := range count {
		ans += count[i]
	}
	fmt.Println(count)
	fmt.Println(ans)
	return ans
}

func inter(pre [][]int, l, r int) int {
	var count int
	for i := 0; i < 26; i++ {
		if pre[r][i] - pre[l][i] > 0 {
			count++
		}
	}
	return count
}

func main() {
	countPalindromicSubsequence("aabca")		// 3
	countPalindromicSubsequence("adc")		// 0
	countPalindromicSubsequence("bbcbaba")	// 4
	countPalindromicSubsequence("ckafnafqo") // 4
}

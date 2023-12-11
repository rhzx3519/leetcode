package main

import "fmt"

/*
*
https://leetcode.cn/problems/next-greater-numerically-balanced-number/?envType=daily-question&envId=2023-12-09
*/
func nextBeautifulNumber(n int) int {
out:
	for n++; ; n++ {
		cnt := [10]int{}
		for x := n; x != 0; x /= 10 {
			cnt[x%10]++
		}
		for x := n; x != 0; x /= 10 {
			if cnt[x%10] != x%10 {
				goto out
			}
		}
		return n
	}
}

func main() {
	fmt.Println(nextBeautifulNumber(1))
}

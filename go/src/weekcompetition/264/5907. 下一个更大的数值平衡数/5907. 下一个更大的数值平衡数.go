/**
@author ZhengHao Lou
Date    2021/10/24
*/
package main

import "fmt"

/**
https://leetcode-cn.com/contest/weekly-contest-264/problems/next-greater-numerically-balanced-number/
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

	return 0
}

func main() {
	fmt.Println(nextBeautifulNumber(1))
	fmt.Println(nextBeautifulNumber(1000))
	fmt.Println(nextBeautifulNumber(3000))
}



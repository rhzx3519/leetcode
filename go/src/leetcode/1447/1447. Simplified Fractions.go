/**
@author ZhengHao Lou
Date    2022/02/10
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/simplified-fractions/
*/
func simplifiedFractions(n int) []string {
	var ans []string
	var counter = map[string]int{}
	for i := 1; i <= n; i++ {
		for j := 1; j < i; j++ {
			g := gcd(i, j)
			s := fmt.Sprintf("%v/%v", j/g, i/g)
			counter[s]++
		}
	}
	for s := range counter {
		ans = append(ans, s)
	}
	fmt.Println(ans)
	return ans
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	simplifiedFractions(4)
}

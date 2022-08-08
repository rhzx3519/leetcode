/**
@author ZhengHao Lou
Date    2021/10/30
*/
package main

import "fmt"

/**
https://leetcode-cn.com/contest/biweekly-contest-64/problems/plates-between-candles/
 */
func platesBetweenCandles(s string, queries [][]int) []int {
	//n := len(s)
	candles := []int{}
	for i, c := range s {
		if c == '|' {
			candles = append(candles, i)
		}
	}

	ans := []int{}
	for _, q := range queries {
		l, r := q[0], q[1]
		fmt.Println(s[l:r+1])
		i, j := lowerBound(candles, l), upperBound(candles,  r)
		if i == j {
			ans = append(ans, 0)
		} else {
			c := j - i
			total := candles[j-1] - candles[i] + 1
			ans = append(ans, total - c)
		}

	}
	return ans
}

func lowerBound(a []int, x int) int {
	l, r := 0, len(a)
	var m int
	for l < r {
		m = l + (r - l)>>1
		if a[m] >= x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func upperBound(a []int, x int) int {
	l, r := 0, len(a)
	var m int
	for l < r {
		m = l + (r - l)>>1
		if a[m] > x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func main() {
	fmt.Println(platesBetweenCandles("**|**|***|", [][]int{{2,5},{5,9}}))
	fmt.Println(platesBetweenCandles("***|**|*****|**||**|*", [][]int{{1,17},{4,5},{14,17},{5,11},{15,16}}))
}

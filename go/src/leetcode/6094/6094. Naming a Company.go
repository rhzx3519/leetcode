/**
@author ZhengHao Lou
Date    2022/06/13
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/naming-a-company/
思路：O(n)
*/
const N = 26

func distinctNames(ideas []string) int64 {
	can := ['z' + 1]['z' + 1]int{}    //  can[i][j] indicates
	cannot := ['z' + 1]['z' + 1]int{} //  can[i][j] indicates
	has := make(map[string]bool)
	for _, idea := range ideas {
		has[idea] = true
	}

	counter := make(map[byte]int)
	for _, idea := range ideas {
		for b := 'a'; b <= 'z'; b++ {
			if !has[string(b)+idea[1:]] {
				can[idea[0]][b]++
			} else {
				cannot[idea[0]][b]++
			}
		}
		counter[idea[0]]++
	}
	var s int64
	for c1 := 'a'; c1 <= 'z'; c1++ {
		for c2 := 'a'; c2 <= 'z'; c2++ {
			if counter[byte(c2)] == 0 || c1 == c2 {
				continue
			}
			s += int64(can[c1][c2] * (counter[byte(c2)] - cannot[c2][c1]))
		}
	}
	fmt.Println(s)
	return s

}

func main() {
	distinctNames([]string{"coffee", "donuts", "time", "toffee"})
	distinctNames([]string{"lack", "back"})
}

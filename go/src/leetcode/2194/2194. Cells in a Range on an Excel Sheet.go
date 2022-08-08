/**
@author ZhengHao Lou
Date    2022/03/08
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/cells-in-a-range-on-an-excel-sheet/
*/
func cellsInRange(s string) []string {
	var ans []string
	c1, r1 := s[0], int(s[1]-'0')
	c2, r2 := s[3], int(s[4]-'0')
	for j := int(c1 - 'A'); j <= int(c2-'A'); j++ {
		for i := r1; i <= r2; i++ {
			ans = append(ans, fmt.Sprintf("%c%v", byte('A'+j), i))
		}
	}
	return ans
}

func main() {
	fmt.Println(cellsInRange("K1:L2"))
}

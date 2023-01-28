/**
@author ZhengHao Lou
Date    2023/01/22
*/
package main

import (
	"fmt"
	"strings"
)

/**
https://leetcode.cn/contest/weekly-contest-329/problems/apply-bitwise-operations-to-make-strings-equal/
思路：枚举
0|1, 0^1 = (1,1) 消除0
1|0, 1^0 = (1,1) 消除0

0|0, 0^0 = (0,0)
1|1, 1^1 = (1,0) 消除1
*/
func makeStringsEqual(s, target string) bool {
	return strings.Contains(s, "1") == strings.Contains(target, "1")
}

func main() {
	//fmt.Println(makeStringsEqual("1010", "0110"))
	//fmt.Println(makeStringsEqual("11", "00"))
	//fmt.Println(makeStringsEqual("101110100", "110011000")) // true
	fmt.Println(makeStringsEqual("001000", "000100")) // true
	
}
